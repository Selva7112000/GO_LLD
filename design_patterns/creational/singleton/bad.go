package singleton

import "fmt"

type BadLogger struct{}

var binstance *BadLogger // not safe for concurrent access

func BGetLogger() *BadLogger {
	if binstance == nil {
		fmt.Println("Creating a new instance of Logger")
		binstance = &BadLogger{}
	}
	return binstance
}

func (l *BadLogger) BLog(message string) {
	fmt.Println("Log:", message)
}

func BadSingleton() {
	logger1 := BGetLogger()
	logger2 := BGetLogger()

	logger1.BLog("Hello from logger1")
	logger2.BLog("Hello from logger2")

	fmt.Println("Same instance?", logger1 == logger2) // true (but not safe in concurrent environments)
}
