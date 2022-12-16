package funcy

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Contains returns true of the slice it's passed contains the
// element it's passed.
func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}

	return false
}

// Reverse returns the slice it's passed in reverse order.
func Reverse[E any](s []E) []E {
	result := make([]E, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return result
}

// Dedupe returns a version of the slice it's passed with
// duplicate elements removed.
func Dedupe[E comparable](s []E) []E {
	result := make([]E, 0, len(s))

	for _, item := range s {
		if !Contains(result, item) {
			result = append(result, item)
		}
	}

	return result
}

// Sort returns a version of the slice it's passed with all
// elements in ascending sorted order.
func Sort[E constraints.Ordered](s []E) []E {
	result := make([]E, len(s))

	copy(result, s)

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

type mapFunc[E any] func(E) E

// Map iterates over the slice it's passed and applies the
// function it's passed to each element, returning a new,
// resulting slice.
func Map[E any](s []E, f mapFunc[E]) []E {
	result := make([]E, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

type keepFunc[E any] func(E) bool

// Filter iterates over the slice it's passed and applies the
// filtering function it's passed to filter out unwanted elements,
// returning a new slice.
func Filter[E any](s []E, f keepFunc[E]) []E {
	result := []E{}

	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

type reduceFunc[E any] func(E, E) E

// Reduce executes the function it's passed on each element of the
// slice it's passed, in order and beginning with the specified initial
// value, passing in the return value from the calculation on the preceding
// element. It returns the final result, after operating on all elements in
// the slice.
func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init

	for _, v := range s {
		cur = f(cur, v)
	}

	return cur
}
