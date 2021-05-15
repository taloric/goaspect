package goaspect

type AtomicFunc func()

type AtomicAct func(AtomicFunc)

type AtomicRet func() (obj interface{})

type AspectFunc interface {

	//combine functions (combine will not do anything , it just combine functions)
	Combine(action AtomicAct) AspectFunc

	//actually do something by the methods below

	//final execute something
	Execute(action AtomicFunc)

	//wrap an object and return it after execute something (it will be more reasonable if golang have generic-type)
	Complete(action AtomicRet) (obj interface{})

	//go with generic-type like
	//Complete[T](f func() T)
}
