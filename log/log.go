package log

import (
	"bytes"
	"errors"
	"github.com/mgutz/ansi"
	"log"
	"time"
)

const (
	LOG_LEVEL_DISABLED = 0
	LOG_LEVEL_ERROR    = 1

	LOG_LEVEL_WARN  = 2
	LOG_LEVEL_INFO  = 3
	LOG_LEVEL_DEBUG = 4
)

var currentLogLevel int = 3
var useColorsForLogger bool = false

var redColorizer func(string) string
var yellowColorizer func(string) string

func init() {

	// initialize the colorizing functions
	// it's faster to use the ansi package via closures, to avoid recalculating ANSI code compilation
	// that would happen every time a ansi.Color is called
	redColorizer = ansi.ColorFunc("red+bh")
	yellowColorizer = ansi.ColorFunc("yellow+bh")

}

func LoggerSetLevel(logLevel int) {

	currentLogLevel = logLevel

}

func LoggerUseColors(useColors bool) {
	useColorsForLogger = useColors
}

func Debug(message string) {

	if currentLogLevel < LOG_LEVEL_DEBUG {
		return
	}

	log.Println("DEBUG: " + message)

}

func DebugNoPrefix(message string) {

	previousFlags := log.Flags()
	log.SetFlags(0)
	Debug(message)
	log.SetFlags(previousFlags)

}

func DebugWithInterface(message string, variantObjects ...interface{}) {

	if currentLogLevel < LOG_LEVEL_DEBUG {
		return
	}

	message = "DEBUG: " + message

	log.Printf(message, variantObjects...)

}

func DebugWithInterfaceNoPrefix(message string, variantObjects ...interface{}) {

	previousFlags := log.Flags()
	log.SetFlags(0)
	DebugWithInterface(message, variantObjects...)
	log.SetFlags(previousFlags)

}

func Info(message string) {

	if currentLogLevel < LOG_LEVEL_INFO {
		return
	}

	log.Println("INFO: " + message)

}

func Warn(message string) {

	if currentLogLevel < LOG_LEVEL_WARN {
		return
	}

	if useColorsForLogger {
		message = yellowColorizer("WARN: " + message)
	} else {
		message = "WARN: " + message
	}

	log.Println(message)

}

func Error(message string, recordedError error) {

	if currentLogLevel < LOG_LEVEL_ERROR {
		return
	}

	if useColorsForLogger {
		message = redColorizer("ERROR: " + message)
	} else {
		message = "ERROR: " + message
	}

	if recordedError != nil {
		log.Println(message, "\n\r", recordedError)
		return
	}

	log.Println(message)

}

func ErrorWithInterface(message string, variantObjects ...interface{}) {

	if currentLogLevel < LOG_LEVEL_ERROR {
		return
	}

	if useColorsForLogger {
		message = redColorizer("ERROR: " + message)
	} else {
		message = "ERROR: " + message
	}

	log.Printf(message+"\n\r", variantObjects...)

}

func Fatal(message string) {

	log.Fatal("FATAL: " + message)

}

func FatalErr(message string, recordedError error) {

	if useColorsForLogger {
		message = redColorizer("FATAL: " + message)
	} else {
		message = "FATAL: " + message
	}

	if recordedError != nil {
		log.Fatal(message, "\n\r", recordedError)
		return
	}

	log.Fatal(message)

}

// Helps track the elapsed time from the beginning to the end of a function.
// It must be run at the very beginning of the function, like this:
// defer log.TimeTrack(time.Now(), "YourFunctionName")
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func Now() time.Time {
	return time.Now()
}

/* BEGIN Error and Logging utility functions */

// Wraps an already existing error with a localized prefix
func NewError(errorPrefix string, originalError error) error {
	return errors.New("ERROR - " + errorPrefix + ": " + originalError.Error())
}

// Wraps local issues in an error format, without needing an already existing error
func NewErrorLocal(errorPrefix string, localError string) error {
	return errors.New("ERROR - " + errorPrefix + ": " + localError)
}

/* BEGIN Journal functions */

type MemoryJournal struct {
	buffer bytes.Buffer
}

func NewMemoryJournal() *MemoryJournal {

	newJournal := &MemoryJournal{}

	return newJournal
}

func (mj *MemoryJournal) AppendLine(entry string) {

	mj.buffer.WriteString(entry)
	mj.buffer.WriteString("\n")
}

func (mj *MemoryJournal) ToString() string {

	return mj.buffer.String()
}

// The memory journal instance cannot be used anylonger after this
func (mj *MemoryJournal) ToStringAndDispose() string {

	finalState := mj.buffer.String()
	mj = nil
	return finalState
}

// Outputs the memory journal contents to the standard output.
// The memory journal instance cannot be used anylonger after this
func (mj *MemoryJournal) ToStdOutAndDispose() {

	log.Print("--------START JOURNAL----------\n" + mj.buffer.String() + "--------END JOURNAL----------\n")
	mj = nil

}
