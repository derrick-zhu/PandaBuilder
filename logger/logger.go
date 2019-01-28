package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	// LevelVerbose show every thing
	LevelVerbose = iota
	// LevelLog only dev log can be shown
	LevelLog
	// LevelError only error log can ben shown
	LevelError
	// LevelConsole is a special level, it has full priviledge to output.
	LevelConsole
)

// internal global variable in logger
var gLogLevel = LevelLog

// Verbose output information with LevelVerbose level
func Verbose(tmpl string, content ...interface{}) {
	consoleOut(os.Stdout, LevelVerbose, tmpl, content...)
}

// Log output information with LevelLog level
func Log(tmpl string, content ...interface{}) {
	consoleOut(os.Stdout, LevelLog, tmpl, content...)
}

// Error output information with LevelError level
func Error(tmpl string, content ...interface{}) {
	consoleOut(os.Stdout, LevelError, tmpl, content...)
}

// Print output information with LevelConsole level
func Print(tmpl string, content ...interface{}) {
	consoleOut(os.Stdout, LevelConsole, tmpl, content...)
}

// Fatal output information with LevelConsole level
func Fatal(tmpl string, content ...interface{}) {
	consoleOut(os.Stdout, LevelConsole, tmpl, content...)
	log.Fatalf(tmpl, content...)
}

func consoleOut(output io.Writer, lvl int, tmpl string, content ...interface{}) {
	if lvl >= gLogLevel || lvl == LevelConsole {
		str := fmt.Sprintf(tmpl, content...)
		output.Write(bytes.NewBufferString(str).Bytes())
	}
}
