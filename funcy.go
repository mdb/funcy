package funcy

import (
  "golang.org/x/exp/constraints"
  "sort"
)

func Contains[E comparable] (s []E, v E) bool {
  for _, vs := range s {
    if v == vs {
      return true
    }
  }

  return false
}

func Reverse[E any](s []E) []E {
  result := make([]E, 0, len(s))

  for i := len(s) - 1; i >= 0; i-- {
    result = append(result, s[i])
  }

  return result
}

func Sort[E constraints.Ordered](s []E) []E {
  result := make([]E, len(s))

  copy(result, s)

  sort.Slice(result, func(i, j int) bool {
    return result[i] < result[j]
  })

  return result
}

type mapFunc[E any] func(E) E

func Map[E any](s []E, f mapFunc[E]) []E {
  result := make([]E, len(s))

  for i := range s {
    result[i] = f(s[i])
  }

  return result
}

type keepFunc[E any] func(E) bool

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

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init

  for _, v := range s {
  	cur = f(cur, v)
  }

  return cur
}
