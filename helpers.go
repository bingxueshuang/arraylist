package alist

import "golang.org/x/exp/constraints"

// Equal is a helper function for types which support equality
// operator (basically wraps up equality operator). It can
// be passed as parameter eq function to methods Contains
// and Index.
func Equal[T comparable](x, y T) bool {
	return x == y
}

// Less is a helper function for types which support less than
// operator (basically wraps up less than operator). It can be
// passed as parameter less function to function IsSorted.
//
// Less is not safe for use with NaN values of float types or
// a few other edge cases. Consider passing a custom less func
// except for simple use cases.
func Less[T constraints.Ordered](x, y T) bool {
	return x < y
}

// IsSorted reports whether the given slice s is sorted or not
// using the given comparison function less.
//
// less func is required to define a strict weak ordering or the
// function may fail to test properly.
func IsSorted[T any](s []T, less func(x, y T) bool) bool {
	for i := len(s) - 1; i > 0; i-- {
		if less(s[i], s[i-1]) {
			return false
		}
	}
	return true
}
