package goaspect

//target function won't execute if condition returns false
func (aspect *Aspect) With(condition func() bool) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		if condition() {
			af()
		}
	})
}

//target function won't execute until condition returns true
func (aspect *Aspect) Until(condition func() bool) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		for {
			if !condition() {
				continue
			}
			af()
			break
		}
	})
}

//target function keep execution while condition returns true
func (aspect *Aspect) While(condition func() bool) *Aspect {
	return aspect.Combine(func(af AtomicFunc) {
		for {
			if condition() {
				af()
			} else {
				break
			}
		}
	})
}
