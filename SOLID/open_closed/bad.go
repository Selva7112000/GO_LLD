package open_closed

import "fmt"

type BadInvoice struct {
	Amount float32
	Type   string
}

func CalculateDiscount(invoice BadInvoice) float32 {
	if invoice.Type == "regular" {
		return invoice.Amount * 0.1
	} else if invoice.Type == "premium" {
		return invoice.Amount * 0.2
	} else if invoice.Type == "vip" {
		return invoice.Amount * 0.3
	}
	return 0
}

func BadOpenClosed() {
	data := BadInvoice{Amount: 1000, Type: "regular"}
	discount := CalculateDiscount(data)
	fmt.Println("\nDiscount:", discount)
}
