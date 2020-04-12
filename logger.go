package main

import (
	"log"
	"os"
)

type Logger struct {
	file *os.File
}

// CreateLogger Constructor of Logger
func CreateLogger(path string) *Logger {
	l := new(Logger)
	var err error
	l.file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

func (l *Logger) write(msg string) {
	l.file.WriteString(msg)
}
