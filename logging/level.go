package logging

import (
	"errors"
	"strings"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

func (level Level) String() string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	default:
		return "DEBUG"
	}
}

func ConvertLevel(levelStr string) (Level, error) {
	lowCase := strings.ToLower(levelStr)

	switch lowCase {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "critical":
		return CRITICAL, nil
	default:
		return 0, errors.New("wrong log level\nplease use debug, info, warning, error, critical")
	}

}
