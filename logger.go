package goaspect

import "fmt"

//logger interface
type Logger interface {
	//log for debug-level(dev members read only) logs
	Debug(message string)
	//log for into-level logs
	Info(message string)
	//log for error-level logs (for problem tracking)
	Error(message string)
	//log for go-errors
	Errorx(err error, message string)
}

type LoggerImpl struct{}

const (
	DEBUG  = "debug"
	INFO   = "info"
	ERROR  = "error"
	ERRORX = "goerror"
)

/*
default implements for interface-Logger down here
*/

func (logger *LoggerImpl) Debug(message string) {
	fmt.Printf("[%s] %s \n", DEBUG, message)
}

func (logger *LoggerImpl) Info(message string) {
	fmt.Printf("[%s] %s \n", INFO, message)
}

func (logger *LoggerImpl) Error(message string) {
	fmt.Printf("[%s] %s \n", ERROR, message)
}

func (logger *LoggerImpl) Errorx(err error, message string) {
	fmt.Printf("[%s] [%s] %s \n", ERRORX, err.Error(), message)
}
