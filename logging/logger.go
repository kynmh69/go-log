package logging

import "log"

type Logger struct {
	logger   log.Logger
	logLevel int
}

func (l *Logger) Info(msg string) {
	if l.logLevel == INFO {
		l.logger.Println()
	}
}

func (l *Logger) Warning(msg string) {
	if l.logLevel == WARNING {
		l.logger.Print()
	}
}