package log

import (
	"fmt"
	"log"
)

type LogLevel string

const (
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
)

func Log(level LogLevel, format string, v ...interface{}) {
	log.Printf("[%s]: %s", level, fmt.Sprintf(format, v...))
}
