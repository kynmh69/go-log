package logging

import (
	"fmt"
	"log"
	"time"
)

type Logger struct {
	logger   log.Logger
	logLevel Level
}

func (l *Logger) print(level Level, msg string) {
	printStr := ""
	printArr := []string{}

	printArr = append(printArr, l.Now())
	printArr = append(printArr, fmt.Sprintf("[%s]", level.String()))
	printArr = append(printArr, msg)
	for i := 0; i < len(printArr); i++ {
		printStr += printArr[i]
		if i == len(printArr)-1 {
			// 最後にスペースは追加しない
			break
		}
		// スペース追加
		printStr += " "
	}
	l.logger.Println(printStr)
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

func (l *Logger) Now() string {
	time := time.Now()
	return time.String()
}
