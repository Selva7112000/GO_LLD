package adapter

import "fmt"

// Target interface that the client expects
type Printer interface {
	PrintMessage(msg string)
}

// Adaptee: A third-party library or legacy class
type LegacyPrinter struct{}

func (lp *LegacyPrinter) Print(msg string) {
	fmt.Println("Legacy Printer:", msg)
}

// Adapter that makes LegacyPrinter compatible with Printer
type LegacyPrinterAdapter struct {
	Legacy *LegacyPrinter
}

func (a *LegacyPrinterAdapter) PrintMessage(msg string) {
	a.Legacy.Print(msg)
}

// Client code using expected interface
func GoodAdapter() {
	var printer Printer = &LegacyPrinterAdapter{
		Legacy: &LegacyPrinter{},
	}
	printer.PrintMessage("Using Adapter Pattern ✅")
}
