package goaspect

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//trace log info
func (runner *Runner) LogInfo(msg string) AspectFunc {
	return runner.Log(msg, "", INFO)
}

//trace log - log to debug level
func (runner *Runner) LogDebug(msg string) AspectFunc {
	return runner.Log(msg, "", DEBUG)
}

//trace log before func execution (info level)
func (runner *Runner) LogBefore(msg string) AspectFunc {
	return runner.Log(msg, "", INFO)
}

//trace log after func execution (info level)
func (runner *Runner) LogAfter(msg string) AspectFunc {
	return runner.Log("", msg, INFO)
}

//trace log before and after func execution (info level)
func (runner *Runner) LogWrap(beforeMsg string, afterMsg string) AspectFunc {
	return runner.Log(beforeMsg, afterMsg, INFO)
}

//trace log core implement
func (runner *Runner) Log(beforeMsg string, afterMsg string, level LogLevel) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		if len(beforeMsg) > 0 {
			switch level {
			case DEBUG:
				(*runner.Logger).Debug(beforeMsg)
			case INFO:
				(*runner.Logger).Info(beforeMsg)
			case ERROR:
				(*runner.Logger).Error(beforeMsg)
			case ERRORX:
				(*runner.Logger).Errorx(errors.New(beforeMsg), beforeMsg)
			}
		}

		af()

		if len(afterMsg) > 0 {
			switch level {
			case DEBUG:
				(*runner.Logger).Debug(afterMsg)
			case INFO:
				(*runner.Logger).Info(afterMsg)
			case ERROR:
				(*runner.Logger).Error(afterMsg)
			case ERRORX:
				(*runner.Logger).Errorx(errors.New(afterMsg), afterMsg)
			}
		}

	})
}

//trace func execution cost time
func (runner *Runner) Watch() AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		start := time.Now()
		(*runner.Logger).Info("------" + start.Format("2006-01-02 15:04:05") + "------")
		af()
		end := time.Now()
		cost := end.Sub(start).Seconds()
		(*runner.Logger).Info("------" + end.Format("2006-01-02 15:04:05") + "------")
		(*runner.Logger).Info("------" + strconv.FormatFloat(cost, 'E', -1, 64) + "------")
	})
}

//only trace panic and log to error level , the panic will still being throw
func (runner *Runner) TrackPanic() AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)

				errInfo := fmt.Sprint(err)
				(*runner.Logger).Errorx(errors.New(errInfo), string(buf))

				//todo : should modify stacktrace info here?
				panic(err)
			}
		}()
		af()
	})
}
