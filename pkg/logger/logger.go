package logger

import (
	"fmt"
	"time"
)

type logLevel uint8

const (
	info logLevel = iota + 1
	error
)

func (l logLevel) toString() string {
	return [...]string{"INFO", "ERROR"}[l-1]
}

func New() *Logger {
	return &Logger{}
}

type Logger struct{}

func (l *Logger) log(level logLevel, message string) {
	stamp := time.Now().Format(time.DateTime)
	fmt.Printf("%s %s: %s\n", stamp, level.toString(), message)
}

func (l *Logger) Info(message string) {
	l.log(info, message)
}

func (l *Logger) Error(message string) {
	l.log(error, message)
}
