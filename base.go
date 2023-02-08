package alist

// List defines a dynamic array which automatically resizes based on
// number of items present.
// This type is not thread-safe.
type List[T any] []T

// Clone returns a shallow copy of the List.
// This method preserves nil slices.
func (a List[T]) Clone() List[T] {
	// preserve nil slices
	if a == nil {
		return nil
	}
	var s List[T] = make([]T, len(a))
	copy(s, a)
	return s
}

// Contains reports whether the item is present in the List a or not.
// It returns if there is any item in a that satisfies the
// eq function for given item.
func (a List[T]) Contains(item T, eq func(x, y T) bool) bool {
	return a.Index(item, eq) >= 0
}

// Index returns the index of first occurrence of item in List,
// or -1 if item not present. It returns the first index of
// item in a that satisfies eq function for given item, or -1
// if no such item exists in the List.
func (a List[T]) Index(item T, eq func(x, y T) bool) int {
	for i, v := range a {
		if eq(v, item) {
			return i
		}
	}
	return -1
}

// Append adds items at the end of the List, growing the List
// if necessary.
func (a List[T]) Append(items ...T) List[T] {
	s := a.growTo(len(a) + len(items))
	return append(s, items...)
}

// Splice replaces s[i:j] by the given items, and returns the modified slice.
// Splice panics if s[i:j] is not a valid slice of List s.
func (a List[T]) Splice(i, j int, items ...T) List[T] {
	_ = a[i:j] // bounds check
	m := len(items)
	n := len(a[:i]) + m + len(a[j:])
	s := a.growTo(n) // grow if necessary
	if n > len(s) {
		s = s[:n] // expand length of slice
	}
	copy(s[i+m:], s[j:])
	copy(s[i:j], items)
	s = s[:n]
	return s.shrink() // shrink if necessary
}

// growTo increases slice capacity if necessary to accommodate a total of n elements.
// growTo does not alter the length of the slice, only capacity.
func (a List[T]) growTo(n int) List[T] {
	if n < 0 {
		panic("cannot be negative")
	}
	c := cap(a)
	if n <= c {
		return a
	}
	// grow capacity by a factor of 2 until enough space
	for n > c {
		c = c * 2
	}
	// reallocate a new slice and copy contents
	s := make([]T, len(a), c)
	copy(s, a)
	return s
}

// shrink clips the slice to a small capacity if length is much smaller than capacity.
func (a List[T]) shrink() List[T] {
	n := len(a)
	c := cap(a)
	// if length is 25% of capacity or lesser
	if n <= c/4 {
		// clip to a capacity of double of length
		return a[: n : 2*n]
	}
	return a
}
