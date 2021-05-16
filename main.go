package goaspect

var DefaultAspect *Aspect

func init() {
	var runner AspectFunc = &Runner{}
	var logger Logger = &LoggerImpl{}

	DefaultAspect = NewAspect(InjectAspectFunc(&runner), InjectLogger(&logger))
}
