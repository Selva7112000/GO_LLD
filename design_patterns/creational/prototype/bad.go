package prototype

import "fmt"

// Without using a prototype, every instance is created from scratch

type BadDocument struct {
	Title   string
	Content string
}

func BadPrototype() {
	// Instead of cloning, we manually create new instances again and again
	doc1 := &BadDocument{
		Title:   "Doc 1",
		Content: "This is the same content",
	}
	doc2 := &BadDocument{
		Title:   "Doc 2",
		Content: "This is the same content",
	}

	fmt.Println("Doc1:", doc1.Title, "-", doc1.Content)
	fmt.Println("Doc2:", doc2.Title, "-", doc2.Content)
}
