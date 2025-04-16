package singleton

import (
	"fmt"
	"sync"
)

type Logger struct{}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		fmt.Println("Creating a new instance of Logger")
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Log(message string) {
	fmt.Println("Log:", message)
}

func GoodSingleton() {
	logger1 := GetLogger()
	logger2 := GetLogger()

	logger1.Log("This is the first log.")
	logger2.Log("This is the second log.")

	fmt.Println("Same instance?", logger1 == logger2)
}
