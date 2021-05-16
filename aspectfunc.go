package goaspect

type AtomicFunc func()

type AtomicAct func(AtomicFunc)

type AtomicRet func() (obj interface{})

type AspectFunc interface {

	//combine functions (do nothing , just combine functions)
	Combine(action AtomicAct) AspectFunc

	//final execute and do something
	Execute(action AtomicFunc)

	//wrap an object and return it after execute (need a better API after golang have generic-type)
	Complete(action AtomicRet) interface{}

	//go with generic-type like
	//Complete[T](f func() T)
}
