package clite

import (
	"fmt"
)

type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

type PrintLogger struct {
}

func (printLogger PrintLogger) Debugf(msg string, tokens ...interface{}) {
	printLogger.Infof(msg, tokens...)
}

func (printLogger PrintLogger) Infof(msg string, tokens ...interface{}) {
	fmt.Printf(msg, tokens...)
}

func (printLogger PrintLogger) Warnf(msg string, tokens ...interface{}) {
	printLogger.Infof(msg, tokens...)
}

func (printLogger PrintLogger) Errorf(msg string, tokens ...interface{}) {
	printLogger.Infof(msg, tokens...)
}
