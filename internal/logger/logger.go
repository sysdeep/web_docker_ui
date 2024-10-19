package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
https://github.com/sirupsen/logrus


log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")
*/

type Logger struct {
	instance *logrus.Logger
}

func NewLogger() *Logger {

	instance := logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	// instance.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	instance.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	instance.SetLevel(logrus.DebugLevel)

	formatter := &logrus.TextFormatter{
		DisableTimestamp: true,
		// DisableQuote: true,
	}

	instance.SetFormatter(formatter)

	return &Logger{instance}
}

// func SetupLogger(logger *LoggerInterface) {
// 	logger = logger
// }

// func Get() *logrus.Logger {
// 	return instance
// }

//--- override ----------------------------------------------------------------

// Info -
func (l *Logger) Info(args ...interface{}) {
	l.instance.Info(args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.instance.Trace(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.instance.Debug(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.instance.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.instance.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	// Calls os.Exit(1) after logging
	l.instance.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	// Calls panic() after logging
	l.instance.Panic(args...)
}

//--- override ----------------------------------------------------------------
