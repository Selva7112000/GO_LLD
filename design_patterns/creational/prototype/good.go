package prototype

import "fmt"

// Cloner interface defines the Clone method
type Cloner interface {
	Clone() Cloner
}

// Document represents a prototype object
type Document struct {
	Title   string
	Content string
}

// Clone creates a copy of the document
func (d *Document) Clone() Cloner {
	copy := *d // shallow copy (fine for simple structs)
	return &copy
}

func GoodPrototype() {
	// Original prototype
	original := &Document{
		Title:   "Prototype Design Pattern",
		Content: "This is a reusable template.",
	}

	// Clone it
	clone1 := original.Clone().(*Document)
	clone1.Title = "Clone 1"

	clone2 := original.Clone().(*Document)
	clone2.Title = "Clone 2"

	fmt.Println("Original:", original.Title, "-", original.Content)
	fmt.Println("Clone 1:", clone1.Title, "-", clone1.Content)
	fmt.Println("Clone 2:", clone2.Title, "-", clone2.Content)
}
