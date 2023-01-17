[![CI](https://github.com/mdb/funcy/actions/workflows/ci.yaml/badge.svg)](https://github.com/mdb/funcy/actions/workflows/ci.yaml) [![PkgGoDev](https://pkg.go.dev/badge/github.com/mdb/funcy)](https://pkg.go.dev/github.com/mdb/funcy)

# funcy ("funky")

A minimal Go package providing a few functional utilities using Go [generics](https://go.dev/blog/intro-generics).

Nothin' too special; largely just for my own learning and casual use, but perhaps this useful to you too.

## Test

```
make
```

## API

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mdb/funcy)](https://pkg.go.dev/github.com/mdb/funcy)

Summary:

```
package funcy // import "github.com/mdb/funcy"

var ErrNotFound = errors.New("not found")
func Any[E any](s []E, f evaluator[E]) bool
func Contains[E comparable](s []E, v E) bool
func Dedupe[E comparable](s []E) []E
func Filter[E any](s []E, f retainer[E]) []E
func Find[E any](s []E, f finder[E]) (E, error)
func Map[E any](s []E, f mapper[E]) []E
func Quicksort[E constraints.Ordered](arr []E) []E
func Reduce[E any](s []E, init E, f reducer[E]) E
func ReduceRight[E any](s []E, init E, f reducer[E]) E
func Reverse[E any](s []E) []E
func Sort[E constraints.Ordered](s []E) []E
```
