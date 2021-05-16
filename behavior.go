package goaspect

import "time"

//Retry once after 1 seconds
func (aspect *Aspect) RetryOnce() *Aspect {
	return aspect.Retry(1, 1000, "")
}

//Retry once after [retryInterval]
func (aspect *Aspect) RetryAfter(retryInterval int) *Aspect {
	return aspect.Retry(1, retryInterval, "")
}

//Retry once with log [msg] after 1 seconds
func (aspect *Aspect) RetryLog(msg string) *Aspect {
	return aspect.Retry(1, 1000, msg)
}

/*
Retry with options

[retryCount : retry times]

[retryInterval : retry after milliseconds]

[msg : log msg every execution]
*/
func (aspect *Aspect) Retry(retryCount int, retryInterval int, msg string) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		retry(retryCount, retryInterval, msg, &aspect.Logger, af)
	})
}

//real retry function
func retry(retryCount int, retryInterval int, msg string, logger *Logger, action AtomicFunc) {
	for i := 0; i <= retryCount; i++ {
		if len(msg) > 0 {
			(*logger).Info(msg)
		}
		action()
		time.Sleep(time.Duration(retryInterval) * time.Millisecond)
	}
}

//delay [delayMilliseconds] to run
func (aspect *Aspect) Delay(delayMilliseconds int) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
		af()
	})
}

//run async with goroutine
func (aspect *Aspect) RunAsync() *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		go af()
	})
}

//run async and callback with goroutine
func (aspect *Aspect) RuncAsyncResult(callBack AtomicFunc) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		go func() {
			af()
			//todo : maybe other ways
			callBack()
		}()
	})
}

//wrap an object and return
func (aspect *Aspect) Wrap(m interface{}, f func(interface{}) interface{}) interface{} {
	//todo : object type validation
	return aspect.Complete(func() (obj interface{}) {
		return f(m)
	})
}
