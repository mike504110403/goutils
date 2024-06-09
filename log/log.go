package log

import (
	"os"
	"runtime"
)

type ilog interface {
	Set(Config)
	Info(string)
	Debug(string)
	Warning(string)
	Error(string)
	Fatal(string)
}

var instance ilog

type (
	EnvMode string
	LogType string
)

const (
	console LogType = "console"
	debug   EnvMode = "debug"
)

// log 設定
type Config struct {
	EnvMode EnvMode
	LogType LogType
}

// 初始化模組
func Init(cfg Config) {
	switch cfg.LogType {
	case console:
		instance = &Console{}
	default:
		instance = &Console{}
	}

	instance.Set(cfg)
}

func Info(msg string) {
	instance.Info(msg)
}

func Debug(msg string) {
	instance.Debug(msg)
}

func Warning(msg string) {
	instance.Warning(msg)
}

func Error(msg string) {
	stack := make([]byte, 2048)
	s := runtime.Stack(stack, false)
	instance.Error(string(stack[0:s]))
	instance.Error(msg)
}

func Fatal(msg string) {
	stack := make([]byte, 2048)
	s := runtime.Stack(stack, false)
	instance.Fatal(string(stack[0:s]))
	instance.Fatal(msg)
	os.Exit(0)
}
