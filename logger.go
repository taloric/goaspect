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

type LoggerImpl struct {
}

type LogLevel string

const (
	DEBUG  LogLevel = "DEBUG"
	INFO   LogLevel = "INFO"
	ERROR  LogLevel = "ERROR"
	ERRORX LogLevel = "GOERROR"
)

/*
default implements for interface-Logger down here
*/

//log debug level
func (logger *LoggerImpl) Debug(message string) {
	fmt.Printf("[%s] %s \n", DEBUG, message)
}

//log info level
func (logger *LoggerImpl) Info(message string) {
	fmt.Printf("[%s] %s \n", INFO, message)
}

//log error level
func (logger *LoggerImpl) Error(message string) {
	fmt.Printf("[%s] %s \n", ERROR, message)
}

//log error level with golang-error
func (logger *LoggerImpl) Errorx(err error, message string) {
	fmt.Printf("[%s] [%s] \n%s \n", ERRORX, err.Error(), message)
}
