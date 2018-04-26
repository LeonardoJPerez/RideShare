package log

import (
	"os"
	"runtime"
	"strconv"

	"github.com/RideShare-Server/utils"
	"github.com/juju/errors"
	logrus "github.com/sirupsen/logrus"
)

const (
	// AuthTopic :
	AuthTopic = "Authorization"

	// BaseServiceTopic :
	BaseServiceTopic = "Base service"

	// DbConnectionTopic :
	DbConnectionTopic = "Database connection"
	// DbOperationTopic :
	DbOperationTopic = "Database operation"

	// EnvironmentVariableTopic :
	EnvironmentVariableTopic = "Environment variable"
)

// InitLog :
func InitLog() {
	// Log as JSON instead of the default ASCII formatter.

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Log by ENV severity or above, otherwise default to logging INFO level severity.
	envValue := utils.GetEnvVariable(utils.LogLevel, logrus.DebugLevel.String())
	logLevel, err := logrus.ParseLevel(envValue)
	if err == nil {
		logrus.SetLevel(logLevel)
	}
}

// Debug :
func Debug(topic string, o interface{}) {
	getEntry(topic).Debug(o)
}

// Info :
func Info(topic, message string) {
	getEntry(topic).Info(message)
}

// Warn :
func Warn(topic string, err error) {
	getEntry(topic).Warn(errors.Cause(err))
}

// Error :
func Error(topic string, err error) {
	getEntry(topic).Error(errors.Cause(err))
}

// Fatal :
func Fatal(topic string, err error) {
	getEntry(topic).Fatal(errors.Cause(err))
}

// Panic :
func Panic(topic string, err error) {
	getEntry(topic).Panic(errors.Cause(err))
}

func getEntry(topic string) *logrus.Entry {
	f, l, fn := getEvent()
	return logrus.WithFields(logrus.Fields{
		"event": fn,
		"topic": topic,
		"file":  f,
		"line":  strconv.Itoa(l),
	})
}

func getEvent() (string, int, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return frame.File, frame.Line, frame.Function
}
