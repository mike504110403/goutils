package log

import (
	"fmt"
	"time"
)

type Console struct {
	Config Config
}

// todo: 預留透過環境變數控制初始化
func (c *Console) Set(cfg Config) {
	c.Config = cfg
}

func (c *Console) Info(msg string) {
	fmt.Printf("%-9s %-19s %s\n", "[Info]", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (c *Console) Debug(msg string) {
	if c.Config.EnvMode == debug {
		fmt.Printf("%-9s %-19s %s\n", "[Debug]", time.Now().Format("2006-01-02 15:04:05"), msg)
	}
}

func (c *Console) Warning(msg string) {
	fmt.Printf("%-9s %-19s %s\n", "[Warning]", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (c *Console) Error(msg string) {
	fmt.Printf("%-9s %-19s %s\n", "[Error]", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (c *Console) Fatal(msg string) {
	fmt.Printf("%-9s %-19s %s\n", "[Fatal]", time.Now().Format("2006-01-02 15:04:05"), msg)
}
