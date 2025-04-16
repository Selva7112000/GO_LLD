package adapter

import "fmt"

// Client directly uses incompatible class
type BadPrinter struct{}

func (bp *BadPrinter) PrintDirectly(lp *LegacyPrinter, msg string) {
	fmt.Print(msg)
}

func BadAdapter() {
	legacy := &LegacyPrinter{}
	bp := &BadPrinter{}
	bp.PrintDirectly(legacy, "No Adapter Used ❌")
}
