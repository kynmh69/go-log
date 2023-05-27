package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	logger   log.Logger
	LogLevel Level
}

func (l *Logger) InitLogger() {
	l.logger = *log.New(os.Stdout, "", log.Lshortfile)
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
	if l.LogLevel <= DEBUG {
		l.print(DEBUG, msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.LogLevel <= INFO {
		l.print(INFO, msg)
	}
}

func (l *Logger) Warning(msg string) {
	if l.LogLevel <= WARNING {
		l.print(WARNING, msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.LogLevel <= ERROR {
		l.print(ERROR, msg)
	}
}

func (l *Logger) Critical(msg string) {
	if l.LogLevel <= CRITICAL {
		l.print(CRITICAL, msg)
	}
}

func (l *Logger) Now() string {
	time := time.Now()
	return time.String()
}
