package goaspect

type AspectFunc interface {
	//combine functions with functions
	Combine(f func()) AspectFunc
	//execute functions
	Execute(f func())
}
