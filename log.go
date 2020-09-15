package log

import "os"

//Logger interface hides logger realization
type Logger interface {
	Info(...interface{})
	Infow(string, ...interface{})
	Error(...interface{})
}

//FlusherFn function to close the logger and flush all pending records before exiting the app
type FlusherFn func() error

//New creates new logger instance with routineName provided
func New(routineName string) (Logger, FlusherFn) {
	if routineName == "" {
		routineName = "noname"
	}
	zl := newZap(routineName, os.Stdout)
	return zl, zl.Sync
}
