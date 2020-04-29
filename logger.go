package main

import (
	"log"
	"os"
)

var logPath string = "./muscletrainer.log" // サービスのデフォルトではルートディレクトリに作成してしまうので設定変更必要

const timeFormat = "2006/01/02 15:04:05 : "

type MTLogger struct {
	logger *log.Logger
}

// NewLogger Constructor of Logger
func NewLogger(path string) *MTLogger {
	l := new(MTLogger)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	l.logger = log.New(f, "", log.Ldate|log.Ltime)
	return l
}

// WriteStartLog Write Start Log to logfile
func (l *MTLogger) WriteStartLog() {
	l.logger.Println("Muscle Trainer Process Start!")
}

func (l *MTLogger) WriteUnluckyLog() {
	l.logger.Println("Unluckey...")
}

func (l *MTLogger) WriteInvalidTimeLog() {
	l.logger.Println("Invalid Time. Please take a rest.")
}

func (l *MTLogger) WriteVideoPlayedLog(url string) {
	l.logger.Println("Muscle Training Video Played: " + url)
}

func (l *MTLogger) WriteStopLog() {
	l.logger.Println("Muscle Trainer Process is Stopped. See you Again!")
}
