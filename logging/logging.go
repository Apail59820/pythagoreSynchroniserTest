package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

var current = INFO
var l = log.New(os.Stdout, "", log.LstdFlags)

// Entry represente une entree de journal.
type Entry struct {
	Time    time.Time
	Level   Level
	Message string
}

var (
	mu      sync.Mutex
	entries []Entry
	// maxEntries limite la taille du buffer de logs conserve en memoire.
	maxEntries = 1000
)

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

func levelString(lv Level) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func addEntry(lv Level, msg string) {
	mu.Lock()
	defer mu.Unlock()
	entries = append(entries, Entry{Time: time.Now(), Level: lv, Message: msg})
	if len(entries) > maxEntries {
		entries = entries[len(entries)-maxEntries:]
	}
}

func Debugf(format string, args ...interface{}) {
	if current <= DEBUG {
		msg := fmt.Sprintf(format, args...)
		l.Printf("DEBUG %s", msg)
		addEntry(DEBUG, msg)
	}
}

func Infof(format string, args ...interface{}) {
	if current <= INFO {
		msg := fmt.Sprintf(format, args...)
		l.Printf("INFO %s", msg)
		addEntry(INFO, msg)
	}
}

func Warnf(format string, args ...interface{}) {
	if current <= WARN {
		msg := fmt.Sprintf(format, args...)
		l.Printf("WARN %s", msg)
		addEntry(WARN, msg)
	}
}

func Errorf(format string, args ...interface{}) {
	if current <= ERROR {
		msg := fmt.Sprintf(format, args...)
		l.Printf("ERROR %s", msg)
		addEntry(ERROR, msg)
	}
}

func Fatalln(args ...interface{}) {
	l.Fatalln(args...)
}

// Entries renvoie la liste des entrees de journal avec un niveau minimum.
func Entries(min Level) []Entry {
	mu.Lock()
	defer mu.Unlock()
	out := make([]Entry, 0, len(entries))
	for _, e := range entries {
		if e.Level >= min {
			out = append(out, e)
		}
	}
	return out
}

// LevelFromString convertit une chaine en niveau de log.
func LevelFromString(s string) Level {
	switch strings.ToLower(s) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	default:
		return DEBUG
	}
}
