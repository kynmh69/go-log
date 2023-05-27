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

// Logger initialize any writer.
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

	printArr = append(printArr, l.now(l.LogFormat))
	printArr = append(printArr, fmt.Sprintf("[%8s]", level.String()))
	printArr = append(printArr, msg)

	for i, str := range printArr {
		printStr += str
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

func (l *Logger) now(format string) string {
	t := time.Now()

	if format == "" {
		format = fmt.Sprintf("%s.000", time.DateTime)
	}

	return t.Format(format)
}
