package goaspect

type Runner struct {
}

//interface implement validation
var _ AspectFunc = (*Runner)(nil)

var ActionChain AtomicAct

//Get a new instance of Runner
func Start() *Runner {
	return &Runner{}
}

func (runner *Runner) Combine(action AtomicAct) AspectFunc {
	if ActionChain == nil {
		ActionChain = action
	} else {
		oldChain := ActionChain
		newChain := func(fun AtomicFunc) {
			//wrap oldchain to chain call ( action1 -> action2 ->..-> real action)
			oldChain(func() {
				//closure here
				action(fun)
			})
		}
		ActionChain = newChain
	}
	return runner
}

func (runner *Runner) Execute(action AtomicFunc) {
	if ActionChain == nil {
		action()
	} else {
		ActionChain(action)
	}
}

func (runner *Runner) Complete(action AtomicRet) (obj interface{}) {
	actionRetVal := action
	if ActionChain == nil {
		return actionRetVal()
	} else {
		//todo : object type validation & object type conversion
		var newObj interface{}
		ActionChain(func() {
			//maybe panic here
			newObj = actionRetVal()
		})
		return newObj
	}
}
