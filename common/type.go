package common

type (
	// Elem is a type of sequence's elements.
	Elem = interface{}

	// Slice is a slice of Elems.
	Slice = []Elem

	// Predicate projects Elem into bool.
	Predicate = func(Elem) bool

	// PredicateIdx projects Elem and its index into bool.
	PredicateIdx = func(Elem, int) bool

	// Comparison compares two Elems and returns negative if the first one is less than the second,
	// zero if the first one is equal to the second and positive if the first one is greater than the second
	// (see https://docs.microsoft.com/dotnet/api/system.comparison-1).
	Comparison = func(Elem, Elem) int

	// Equality determines whether two Elems are equal.
	Equality = func(Elem, Elem) bool

	// Less determines whether the first Elem is less than the second.
	Less = func(Elem, Elem) bool
)
