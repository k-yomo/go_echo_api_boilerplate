package logger

import (
	"bou.ke/monkey"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"testing"
)

type testBuffer struct {
	*bytes.Buffer
}

type jsonLog struct {
	Level   string            `json:"level"`
	Msg     string            `json:"msg"`
	JsonMsg map[string]string `json:"jsonMsg"`
}

func unmartialJsonLog(j []byte) *jsonLog {
	log := new(jsonLog)
	json.Unmarshal(j, &log)
	return log
}

func (testBuffer) Sync() error {
	return nil
}

func newTestLogger() (echo.Logger, *bytes.Buffer) {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	b := new(bytes.Buffer)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), testBuffer{Buffer: b}, zap.DebugLevel)
	return NewLogger(zap.New(core)), b
}

func TestOutput(t *testing.T) {
	l, _ := newTestLogger()
	out := l.Output()
	size, err := out.Write(make([]byte, 10))
	assert.Equal(t, ioutil.Discard, out)
	assert.Equal(t, 10, size)
	assert.Equal(t, nil, err)
}

func TestCustomize(t *testing.T) {
	l, _ := newTestLogger()
	l.SetLevel(log.ERROR)  // do nothing
	l.SetOutput(os.Stdout) // do nothing
	l.SetPrefix("prefix")  // do nothing
	l.SetHeader("header")  // do nothing

	assert.Equal(t, "", l.Prefix())
	assert.Equal(t, log.DEBUG, l.Level())
}

func TestPrint(t *testing.T) {
	l, b := newTestLogger()
	l.Print("test")
	assert.Contains(t, b.String(), "info")
	assert.Contains(t, b.String(), "test")
}

func TestPrintf(t *testing.T) {
	l, b := newTestLogger()
	l.Printf("test %d", 1)
	assert.Contains(t, b.String(), "test 1")
}

func TestPrintj(t *testing.T) {
	l, b := newTestLogger()
	l.Printj(log.JSON{"key": "test"})
	assert.Equal(t, &jsonLog{"info", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}

func TestDebug(t *testing.T) {
	l, b := newTestLogger()
	l.Debug("test")
	assert.Contains(t, b.String(), "debug")
	assert.Contains(t, b.String(), "test")
}

func TestDebugf(t *testing.T) {
	l, b := newTestLogger()
	l.Debugf("test %d", 1)
	assert.Contains(t, b.String(), "debug")
	assert.Contains(t, b.String(), "test 1")
}

func TestDebugj(t *testing.T) {
	l, b := newTestLogger()
	l.Debugj(log.JSON{"key": "test"})
	assert.Equal(t, &jsonLog{"debug", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}

func TestInfo(t *testing.T) {
	l, b := newTestLogger()
	l.Info("test")
	assert.Contains(t, b.String(), "info")
	assert.Contains(t, b.String(), "test")
}

func TestInfof(t *testing.T) {
	l, b := newTestLogger()
	l.Infof("test %d", 1)
	assert.Contains(t, b.String(), "info")
	assert.Contains(t, b.String(), "test 1")
}

func TestInfoj(t *testing.T) {
	l, b := newTestLogger()
	l.Infoj(log.JSON{"key": "test"})
	assert.Equal(t, &jsonLog{"info", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}

func TestWarn(t *testing.T) {
	l, b := newTestLogger()
	l.Warn("test")
	assert.Contains(t, b.String(), "warn")
	assert.Contains(t, b.String(), "test")
}

func TestWarnf(t *testing.T) {
	l, b := newTestLogger()
	l.Warnf("test %d", 1)
	assert.Contains(t, b.String(), "warn")
	assert.Contains(t, b.String(), "test 1")
}

func TestWarnj(t *testing.T) {
	l, b := newTestLogger()
	l.Warnj(log.JSON{"key": "test"})
	assert.Equal(t, &jsonLog{"warn", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}

func TestError(t *testing.T) {
	l, b := newTestLogger()
	l.Error("test")
	assert.Contains(t, b.String(), "error")
	assert.Contains(t, b.String(), "test")
}

func TestErrorf(t *testing.T) {
	l, b := newTestLogger()
	l.Errorf("test %d", 1)
	assert.Contains(t, b.String(), "error")
	assert.Contains(t, b.String(), "test 1")
}

func TestErrorj(t *testing.T) {
	l, b := newTestLogger()
	l.Errorj(log.JSON{"key": "test"})
	assert.Equal(t, &jsonLog{"error", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}

func TestFatal(t *testing.T) {
	exitCode := 0
	monkey.Patch(os.Exit, func(code int) { exitCode = code })
	l, b := newTestLogger()
	l.Fatal("test")
	assert.NotEqual(t, 2, exitCode)
	assert.Contains(t, b.String(), "fatal")
	assert.Contains(t, b.String(), "test")
}

func TestFatalf(t *testing.T) {
	exitCode := 0
	monkey.Patch(os.Exit, func(code int) { exitCode = code })
	l, b := newTestLogger()
	l.Fatalf("test %d", 1)
	assert.NotEqual(t, 2, exitCode)
	assert.Contains(t, b.String(), "fatal")
	assert.Contains(t, b.String(), "test 1")
}

func TestFatalj(t *testing.T) {
	exitCode := 0
	monkey.Patch(os.Exit, func(code int) { exitCode = code })
	l, b := newTestLogger()
	l.Fatalj(log.JSON{"key": "test"})
	assert.NotEqual(t, 2, exitCode)
	assert.Equal(t, &jsonLog{"fatal", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}

func TestPanic(t *testing.T) {
	exitCode := 0
	monkey.Patch(os.Exit, func(code int) { exitCode = code })
	l, b := newTestLogger()
	assert.Panics(t, func() { l.Panic("test") })
	assert.NotEqual(t, 2, exitCode)
	assert.Contains(t, b.String(), "panic")
	assert.Contains(t, b.String(), "test")
}

func TestPanicf(t *testing.T) {
	exitCode := 0
	monkey.Patch(os.Exit, func(code int) { exitCode = code })
	l, b := newTestLogger()
	assert.Panics(t, func() { l.Panicf("test %d", 1) })
	assert.NotEqual(t, 2, exitCode)
	assert.Contains(t, b.String(), "panic")
	assert.Contains(t, b.String(), "test 1")
}

func TestPanicj(t *testing.T) {
	exitCode := 0
	monkey.Patch(os.Exit, func(code int) { exitCode = code })
	l, b := newTestLogger()
	assert.Panics(t, func() { l.Panicj(log.JSON{"key": "test"}) })
	assert.NotEqual(t, 2, exitCode)
	assert.Equal(t, &jsonLog{"panic", "echo json log", map[string]string{"key": "test"}}, unmartialJsonLog(b.Bytes()))
}
