package goaspect

type Aspect struct {
	AspectFunc
	Logger
}

type AspectInject func(*Aspect)

//Get a new instance of Aspect
func NewAspect(injects ...AspectInject) *Aspect {
	aspect := &Aspect{}
	for _, i := range injects {
		i(aspect)
	}
	return aspect
}

func InjectAspectFunc(aspectFunc *AspectFunc) AspectInject {
	return func(a *Aspect) {
		a.AspectFunc = *aspectFunc
	}
}

func InjectLogger(logger *Logger) AspectInject {
	return func(a *Aspect) {
		a.Logger = *logger
	}
}

//this method is for chain methods
func (aspect *Aspect) Combine(action AtomicAct) *Aspect {
	if aspect.AspectFunc != nil {
		//real combination
		aspect.AspectFunc.Combine(action)
	}
	return aspect
}
