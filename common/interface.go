package common

// Counter is the interface that wraps the Count method.
type Counter interface {
	// Count returns the number of elements in a sequence.
	Count() int
}

// Itemer is the interface that wraps the Item method.
type Itemer interface {
	// Item returns the element of a sequence at the specified index.
	Item(int) Elem
}
