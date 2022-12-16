package funcy

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Contains returns true if the slice it's passed contains the
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

// Dedupe returns the slice it's passed with duplicate elements removed.
func Dedupe[E comparable](s []E) []E {
	result := make([]E, 0, len(s))

	for _, item := range s {
		if !Contains(result, item) {
			result = append(result, item)
		}
	}

	return result
}

func partition[E constraints.Ordered](arr []E, low, high int) ([]E, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func quicksort[E constraints.Ordered](arr []E, low, high int) []E {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quicksort(arr, low, p-1)
		arr = quicksort(arr, p+1, high)
	}

	return arr
}

// Quicksort implements the quicksort algorithm to sort the slice
// it's passed.
func Quicksort[E constraints.Ordered](arr []E) []E {
	return quicksort(arr, 0, len(arr)-1)
}

// Sort returns the slice it's passed with all elements in ascending
// sorted order. It uses Quicksort.
func Sort[E constraints.Ordered](s []E) []E {
	return Quicksort(s)
}

type mapper[E any] func(E) E

// Map applies the function it's passed to each element in the slice
// it's passed and returns the resulting slice.
func Map[E any](s []E, f mapper[E]) []E {
	result := make([]E, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

type retainer[E any] func(E) bool

// Filter applies the function it's passed to each element in the slice
// it's passed to filter out unwanted elements and return the resulting,
// filtered slice.
func Filter[E any](s []E, f retainer[E]) []E {
	result := []E{}

	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

type finder[E any] func(E) bool

// Find applies the function it's passed to each element in the slice
// and returns the first element for which the function returns true.
func Find[E any](s []E, f finder[E]) (E, error) {
	for _, v := range s {
		if f(v) {
			return v, nil
		}
	}

	var e E

	return e, errors.New("not found")
}

type evaluator[E any] func(E) bool

// Any applies the function it's passed to each element in the slice
// and returns true for the first element for which the function returns
// true. Otherwise, it returns false.
func Any[E any](s []E, f evaluator[E]) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}

	return false
}

type reducer[E any] func(E, E) E

// Reduce executes the function it's passed on each element of the
// slice it's passed, in order and beginning with the specified initial
// value, passing in the return value from the calculation on the preceding
// element. It returns the final result, after operating on all elements in
// the slice.
func Reduce[E any](s []E, init E, f reducer[E]) E {
	cur := init

	for _, v := range s {
		cur = f(cur, v)
	}

	return cur
}
