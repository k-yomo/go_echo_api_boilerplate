package logger

import (
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
)

// echoLogger struct wraps zap logger
type echoLogger struct {
	l *zap.SugaredLogger
}

// NewLogger converts zap.Logger to echo.Logger
func NewLogger(logger *zap.Logger) *echoLogger {
	return &echoLogger{
		l: logger.
			// for skipping wrapping function
			WithOptions(zap.AddCallerSkip(1)).
			Sugar(),
	}
}

// Output writer
func (e *echoLogger) Output() io.Writer {
	return ioutil.Discard
}

// SetOutput do nothing
func (e *echoLogger) SetOutput(w io.Writer) {}

// SetHeader do nothing
func (e *echoLogger) SetHeader(h string) {}

// Prefix do nothing
func (e *echoLogger) Prefix() string {
	return ""
}

// SetPrefix do nothing
func (e *echoLogger) SetPrefix(p string) {}

// Level do nothing
func (e *echoLogger) Level() log.Lvl {
	return log.DEBUG
}

// SetLevel do nothing
func (e *echoLogger) SetLevel(v log.Lvl) {}

// Print uses fmt.Sprint to construct and log a message.
func (e *echoLogger) Print(i ...interface{}) {
	e.l.Info(i...)
}

// Printf uses fmt.Sprintf to log a templated message.
func (e *echoLogger) Printf(format string, args ...interface{}) {
	e.l.Infof(format, args...)
}

// Printj logs a message with some additional custom_context.
func (e *echoLogger) Printj(j log.JSON) {
	e.l.Infow("echo json log", "jsonMsg", j)
}

// Debug uses fmt.Sprint to construct and log a message.
func (e *echoLogger) Debug(i ...interface{}) {
	e.l.Debug(i...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (e *echoLogger) Debugf(format string, args ...interface{}) {
	e.l.Debugf(format, args...)
}

// Debugj logs a message with some additional custom_context.
func (e *echoLogger) Debugj(j log.JSON) {
	e.l.Debugw("echo json log", "jsonMsg", j)
}

// Info uses fmt.Sprint to construct and log a message.
func (e *echoLogger) Info(i ...interface{}) {
	e.l.Infow("echo log", i...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (e *echoLogger) Infof(format string, args ...interface{}) {
	e.l.Infof(format, args...)
}

// Infoj logs a message with some additional custom_context.
func (e *echoLogger) Infoj(j log.JSON) {
	e.l.Infow("echo json log", "jsonMsg", j)
}

// Warn uses fmt.Sprint to construct and log a message.
func (e *echoLogger) Warn(i ...interface{}) {
	e.l.Warnw("echo log", i...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (e *echoLogger) Warnf(format string, args ...interface{}) {
	e.l.Warnf(format, args...)
}

// Warnj logs a message with some additional custom_context.
func (e *echoLogger) Warnj(j log.JSON) {
	e.l.Warnw("echo json log", "jsonMsg", j)
}

// Error uses fmt.Sprint to construct and log a message.
func (e *echoLogger) Error(i ...interface{}) {
	e.l.Errorw("echo log", i...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (e *echoLogger) Errorf(format string, args ...interface{}) {
	e.l.Errorf(format, args...)
}

// Errorj logs a message with some additional custom_context.
func (e *echoLogger) Errorj(j log.JSON) {
	e.l.Errorw("echo json log", "jsonMsg", j)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (e *echoLogger) Fatal(i ...interface{}) {
	e.l.Fatalw("echo log", i...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (e *echoLogger) Fatalf(format string, args ...interface{}) {
	e.l.Fatalf(format, args...)
}

// Fatalj logs a message with some additional custom_context, then calls os.Exit.
func (e *echoLogger) Fatalj(j log.JSON) {
	e.l.Fatalw("echo json log", "jsonMsg", j)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (e *echoLogger) Panic(i ...interface{}) {
	e.l.Panicw("echo log", i...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (e *echoLogger) Panicf(format string, args ...interface{}) {
	e.l.Panicf(format, args...)
}

// Panicj logs a message with some additional custom_context, then panics.
func (e *echoLogger) Panicj(j log.JSON) {
	e.l.Panicw("echo json log", "jsonMsg", j)
}
