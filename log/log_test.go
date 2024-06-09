package log

import (
	"os"
	"testing"
)

func Set() {
	os.Setenv("Env_Mode", "debug")
	os.Setenv("Log_Type", "console")
}

func TestConsole(t *testing.T) {
	Set()
	Init(Config{
		LogType: LogType(os.Getenv("Log_Type")),
		EnvMode: EnvMode(os.Getenv("Env_Mode")),
	})
	//Info("msg Info")
	Debug("msg Debug")
	//Warning("msg Warning")
	//Error("msg Error")
	//Fatal("msg Fatal")
	os.Setenv("Env_Mode", "online")
	Init(Config{
		LogType: LogType(os.Getenv("Log_Type")),
		EnvMode: EnvMode(os.Getenv("Env_Mode")),
	})
	Debug("msg Deug")
}
