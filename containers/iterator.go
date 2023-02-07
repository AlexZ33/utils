package containers

//IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
type IteratorWithIndex interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
	Value() interface{}

	// Index returns the current element's index.
	// Does not modify the state of the iterator.
	Index() int

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
	Begin()
	// First moves the iterator to the first element and returns true if there was a first element in the container
	First() bool

	// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	NextTo(func(index int, value interface{}) bool) bool
}
