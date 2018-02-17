// log is designed like github.com/brutella to make logs look similar
package log

import (
	"log"
	"os"
)

var (
	// Debug is a logger with a DEBUG prefix
	Debug = log.New(os.Stdout, "DEBUG ", log.LstdFlags)
	// Info is a logger with an INFO prefix
	Info = log.New(os.Stdout, "INFO ", log.LstdFlags)
	// Error is a logger with an ERROR prefix
	Error = log.New(os.Stdout, "ERROR ", log.LstdFlags)
	// Warn is a logger with a WARN prefix
	Warn = log.New(os.Stdout, "WARN ", log.LstdFlags)
)
