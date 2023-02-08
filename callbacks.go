package alist

// Every iterates over each item of the slice s and calls the given function f.
// Every returns true if the function f returns true for every item of s.
func Every[T any](s []T, f func(i int, v T) bool) bool {
	for i, v := range s {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Some iterates over each item of the slice s and calls the given function f.
// Some returns true if the function f returns true for some item in s.
func Some[T any](s []T, f func(i int, v T) bool) bool {
	for i, v := range s {
		if f(i, v) {
			return true
		}
	}
	return false
}

// Filter returns a new slice containing all items of s for which the given
// function f returns a true value.
func Filter[T any](s []T, f func(i int, v T) bool) []T {
	if s == nil {
		return nil
	}
	// keep enough capacity to avoid prevent re-alloc
	a := make([]T, 0, len(s))
	for i, v := range s {
		if f(i, v) {
			a = append(a, v)
		}
	}
	// clip the capacity
	return a[:len(a):len(a)]
}

// Reduce executes a given "reducer" function f on each item of the slice,
// in order along with result of f on the preceding item of s. Final value
// obtained by "reducing" all elements of s to a single value is returned.
// For the first iteration, initial value is passed as "prev" value for the
// reducer function.
func Reduce[T, U any](s []T, f func(i int, v T, prev U) U, initial U) U {
	state := initial
	for i, v := range s {
		state = f(i, v, state)
	}
	return state
}

// ReduceSimple is similar to Reduce except that it "reduces" the items of s
// taking the first item itself as initial value for the reducer function.
// Final value obtained by "reducing" all elements of s is the same type as
// its members.
// ReduceSimple panics if the slice is empty (since it needs the first item
// as the initial value)
func ReduceSimple[T any](s []T, f func(i int, v T, prev T) T) T {
	if len(s) == 0 {
		panic("empty slice")
	}
	return Reduce(s[1:], f, s[0])
}
