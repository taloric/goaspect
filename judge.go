package goaspect

//func not execute if condition returns false
func (runner *Runner) With(condition func() bool) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		if condition() {
			af()
		}
	})
}

//func not execute until condition returns true
func (runner *Runner) Until(condition func() bool) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		for {
			if !condition() {
				continue
			}
			af()
			break
		}
	})
}

//func will execute while condition returns true
func (runner *Runner) While(condition func() bool) AspectFunc {
	return runner.Combine(func(af AtomicFunc) {
		for {
			if condition() {
				af()
			} else {
				break
			}
		}
	})
}
