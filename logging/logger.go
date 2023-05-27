package logging

import (
	"fmt"
	"io"
	"log"
	"time"
)

type Logger struct {
	Logger    []log.Logger
	LogLevel  Level
	LogFormat string
}

func New(streams []io.Writer) *Logger {
	l := &Logger{}
	for _, stream := range streams {
		l.Logger = append(l.Logger, *log.New(stream, "", log.Lshortfile))
	}
	return l
}

func (l *Logger) print(level Level, msg string) {
	printStr := ""
	printArr := []string{}

	printArr = append(printArr, l.Now(l.LogFormat))
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
	for _, stream := range l.Logger {
		stream.Println(printStr)
	}
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

func (l *Logger) Now(format string) string {
	time := time.Now()

	if format == "" {
		format = "2006-01-02 15:04:05"
	}

	return time.Format(format)
}
