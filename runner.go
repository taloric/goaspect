package goaspect

//an implement of Interface AspectFunc
type Runner struct {
	actinoChain AtomicAct
}

func (runner *Runner) Combine(action AtomicAct) AspectFunc {
	if runner.actinoChain == nil {
		runner.actinoChain = action
	} else {
		oldChain := runner.actinoChain
		newChain := func(fun AtomicFunc) {
			//wrap oldchain to chain call ( action1 -> action2 ->..-> real action)
			//actually call chain like : action1( action2 ( action3 ( real action ) ) )
			//an old action wrap new action and make new action as a parameter to execute
			oldChain(func() {
				action(fun)
			})
		}
		runner.actinoChain = newChain
	}
	return runner
}

func (runner *Runner) Execute(action AtomicFunc) {
	if runner.actinoChain == nil {
		action()
	} else {
		//when really execute something, action chain wrap real action as a func parameter
		runner.actinoChain(action)
	}
}

func (runner *Runner) Complete(action AtomicRet) interface{} {
	actionRetVal := action
	if runner.actinoChain == nil {
		return actionRetVal()
	} else {
		//todo : object type validation & object type conversion
		var newObj interface{}
		runner.actinoChain(func() {
			//maybe panic here
			newObj = actionRetVal()
		})
		return newObj
	}
}
