package debug

import (
	"fmt"
	"time"
)

// Levels of debug.
const (
	LevelSilent = iota - 1
	LevelPrint
	LevelLog
	LevelTrace
)

// Level of debug.
var Level = LevelLog

// Print if Level is 0 or more.
func Print(vs ...interface{}) {
	if Level >= LevelPrint {
		fmt.Println(vs...)
	}
}

// Log if Level is 1 or more.
func Log(vs ...interface{}) {
	if Level >= LevelLog {
		fmt.Println(vs...)
	}
}

// Logf if Level is 1 or more.
func Logf(format string, vs ...interface{}) {
	if Level >= LevelLog {
		fmt.Printf(format, vs...)
	}
}

// Trace if Level is 2 or more.
func Trace(vs ...interface{}) {
	if Level >= LevelTrace {
		fmt.Println(vs...)
	}
}

var tick time.Time

//LogT Log throttled
func LogT(dt time.Duration, vs ...interface{}) {
	if time.Since(tick) < dt {
		return
	}
	tick = time.Now()
	Log(vs...)
}

//TraceT Trace throttled
func TraceT(dt time.Duration, vs ...interface{}) {
	if time.Since(tick) < dt {
		return
	}
	tick = time.Now()
	Trace(vs...)
}
