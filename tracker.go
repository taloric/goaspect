package goaspect

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//trace log info
func (aspect *Aspect) LogInfo(msg string) *Aspect {
	return aspect.Log(msg, "", INFO)
}

//trace log - log to debug level
func (aspect *Aspect) LogDebug(msg string) *Aspect {
	return aspect.Log(msg, "", DEBUG)
}

//trace log before func execution (info level)
func (aspect *Aspect) LogBefore(msg string) *Aspect {
	return aspect.Log(msg, "", INFO)
}

//trace log after func execution (info level)
func (aspect *Aspect) LogAfter(msg string) *Aspect {
	return aspect.Log("", msg, INFO)
}

//trace log before and after func execution (info level)
func (aspect *Aspect) LogWrap(beforeMsg string, afterMsg string) *Aspect {
	return aspect.Log(beforeMsg, afterMsg, INFO)
}

//trace log implement
func (aspect *Aspect) Log(beforeMsg string, afterMsg string, level LogLevel) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		if len(beforeMsg) > 0 {
			switch level {
			case DEBUG:
				aspect.Logger.Debug(beforeMsg)
			case INFO:
				aspect.Logger.Info(beforeMsg)
			//should not trace error level log
			case ERROR:
				aspect.Logger.Error(beforeMsg)
			case ERRORX:
				aspect.Logger.Errorx(errors.New(beforeMsg), beforeMsg)
			}
		}

		af()

		if len(afterMsg) > 0 {
			switch level {
			case DEBUG:
				aspect.Logger.Debug(afterMsg)
			case INFO:
				aspect.Logger.Info(afterMsg)
			case ERROR:
				aspect.Logger.Error(afterMsg)
			case ERRORX:
				aspect.Logger.Errorx(errors.New(afterMsg), afterMsg)
			}
		}

	})
}

//trace func execution cost
func (aspect *Aspect) Watch() *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		start := time.Now()
		aspect.Logger.Info("------" + start.Format("2006-01-02 15:04:05") + "]------")
		af()
		end := time.Now()
		cost := end.Sub(start).Seconds()
		aspect.Logger.Info("------[" + end.Format("2006-01-02 15:04:05") + "]------")
		aspect.Logger.Info("------[" + strconv.FormatFloat(cost, 'E', -1, 64) + "]------")
	})
}

//only trace panic and log to error level , the panic will still being throw
func (aspect *Aspect) TrackPanic() *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)

				errInfo := fmt.Sprint(err)
				aspect.Logger.Errorx(errors.New(errInfo), string(buf))

				//todo : should modify stacktrace info here?
				panic(err)
			}
		}()
		af()
	})
}
