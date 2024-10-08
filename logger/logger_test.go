package logger

import (
	"errors"
	"testing"
)

func TestPrint(t *testing.T) {
	logsDir = "./logs"
	Print("hello")
}

func TestWarning(t *testing.T) {
	logsDir = "./logs"
	Warning("hello")
}

func TestError(t *testing.T) {
	logsDir = "./logs"
	Error(errors.New("hello"))
}

func TestDebug(t *testing.T) {
	logsDir = "./logs"
	Debug(0, false, "hello")
	Debug(0, true, "hello2")
	Debug(3, true, "hello3")
	SetDebugLevel(3)
	Debug(3, true, "hello4")
}

func TestInfo(t *testing.T) {
	logsDir = "./logs"
	Info("hello")
}
