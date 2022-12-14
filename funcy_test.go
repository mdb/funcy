package funcy

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	list := []string{"foo", "bar", "baz"}

	got := Contains(list, "foo")
	expected := true
	if got != expected {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

	got = Contains(list, "bim")
	expected = false
	if got != expected {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func TestReverse(t *testing.T) {
	list := []string{"foo", "bar", "baz"}

	got := Reverse(list)
	expected := []string{"baz", "bar", "foo"}
	if strings.Join(got, " ") != strings.Join(expected, " ") {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func TestSort(t *testing.T) {
	list := []int{3, 1, 2}

	got := Reverse(list)
	expected := []int{1, 2, 3}
	if expected[0] != got[0] && expected[2] != got[2] {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func TestMap(t *testing.T) {
	list := []int{3, 1, 2}

	got := Map(list, func(item int) int {
		return item + 2
	})
	expected := []int{5, 3, 4}
	if expected[0] != got[0] && expected[2] != got[2] {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func TestFilter(t *testing.T) {
	list := []int{3, 1, 2}

	got := Filter(list, func(item int) bool {
		return item == 3
	})
	expected := []int{3}
	if expected[0] != got[0] && len(expected) != len(got) {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func TestReduce(t *testing.T) {
	list := []int{3, 1, 2}

	got := Reduce(list, 0, func(cur, val int) int {
		return cur + val
	})
	expected := 6
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

	got = Reduce(list, 2, func(cur, val int) int {
		return cur + val
	})
	expected = 8
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}