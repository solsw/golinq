package common

var (
	// Identity is a selector that projects the Elem into itself.
	Identity = func(el Elem) Elem { return el }
)
