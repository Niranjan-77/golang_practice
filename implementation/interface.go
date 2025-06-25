package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

const format = "%v: Info: %s /n"

// Logger is logging interface
type Logger interface {
	Info(string)
}

// ScreenLogger is log impl
type ScreenLogger struct{}

// Info impl for ScreenLogger
func (ScreenLogger) Info(message string) {
	fmt.Printf(format, time.Now(), message)
}

// FileLogger is log impl
type FileLogger struct {
	File *os.File
}

// Info impl for FileLogger
func (l *FileLogger) Info(message string) {
	fmt.Fprintf(l.File, format, time.Now(), message)
}

// NewFileLogger creates a new FileLogger
func NewFileLogger(filename string) *FileLogger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &FileLogger{file}
}
func main() {
	var log Logger
	log = NewFileLogger("log.txt")
	log.Info("This is a screen log message.")
	fmt.Println(reflect.TypeOf(log))
}
