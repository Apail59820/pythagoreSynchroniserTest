package logging

import (
	"log"
	"os"
	"strings"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var current = INFO
var l = log.New(os.Stdout, "", log.LstdFlags)

// Setup initialise le niveau de journalisation depuis la variable LOG_LEVEL.
func Setup() {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		current = DEBUG
	case "warn":
		current = WARN
	case "error":
		current = ERROR
	default:
		current = INFO
	}
}

func Debugf(format string, args ...interface{}) {
	if current <= DEBUG {
		l.Printf("DEBUG "+format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if current <= INFO {
		l.Printf("INFO "+format, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if current <= WARN {
		l.Printf("WARN "+format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if current <= ERROR {
		l.Printf("ERROR "+format, args...)
	}
}

func Fatalln(args ...interface{}) {
	l.Fatalln(args...)
}
