package goaspect

import "time"

//Retry once after 1 seconds
func (runner *Runner) RetryOnce() AspectFunc {
	return runner.Retry(1, 1000, "")
}

//Retry once after interval
func (runner *Runner) RetryAfter(retryInterval int) AspectFunc {
	return runner.Retry(1, retryInterval, "")
}

//Retry with log every time
func (runner *Runner) RetryLog(msg string) AspectFunc {
	return runner.Retry(1, 1000, msg)
}

//Retry with options
func (runner *Runner) Retry(retryCount int, retryInterval int, msg string) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		retry(retryCount, retryInterval, msg, runner.Logger, af)
	})
}

//delay [delayMilliseconds] to run
func (runner *Runner) Delay(delayMilliseconds int) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
		af()
	})
}

//run async with goroutine
func (runner *Runner) RunAsync() AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		go af()
	})
}

//run callback async with goroutine
func (runner *Runner) RuncAsyncResult(callBack AtomicFunc) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		go func() {
			af()
			callBack()
		}()
	})
}

//retry core func
func retry(retryCount int, retryInterval int, msg string, logger *Logger, action AtomicFunc) {
	for i := 0; i <= retryCount; i++ {
		if len(msg) > 0 {
			(*logger).Info(msg)
		}
		action()
		time.Sleep(time.Duration(retryInterval) * time.Millisecond)
	}
}
