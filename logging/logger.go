package logging

import "log"

type Logger struct {
	logger   log.Logger
	logLevel Level
}

func (l *Logger) print(level Level, msg string) {
	l.logger.Printf("[%s]: %s", level, msg)
}

func (l *Logger) Debug(msg string) {
	if l.logLevel == DEBUG {
		l.print(DEBUG, msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.logLevel == INFO {
		l.print(INFO, msg)
	}
}

func (l *Logger) Warning(msg string) {
	if l.logLevel == WARNING {
		l.print(WARNING, msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.logLevel == ERROR {
		l.print(ERROR, msg)
	}
}

func (l *Logger) Critical(msg string) {
	if l.logLevel == CRITICAL {
		l.logger.Println(CRITICAL, msg)
	}
}
