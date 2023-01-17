package funcy

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleContains() {
	list := []string{"foo", "bar", "baz"}

	fmt.Println(Contains(list, "foo"))
	//
	// Output: true
}

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

func ExampleDedupe() {
	list := []string{"foo", "foo", "bar"}

	fmt.Println(Dedupe(list))
	//
	// Output: [foo bar]
}

func TestDedupe(t *testing.T) {
	list := []string{"foo", "foo", "bar"}

	got := Dedupe(list)
	expected := []string{"foo", "bar"}
	if len(got) != len(expected) || got[0] != expected[0] || got[1] != expected[1] {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func ExampleReverse() {
	list := []string{"foo", "bar", "baz"}

	fmt.Println(Reverse(list))
	//
	// Output: [baz bar foo]
}

func TestReverse(t *testing.T) {
	list := []string{"foo", "bar", "baz"}

	got := Reverse(list)
	expected := []string{"baz", "bar", "foo"}
	if strings.Join(got, " ") != strings.Join(expected, " ") {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func ExampleSort() {
	list := []int{3, 1, 2}

	fmt.Println(Sort(list))
	//
	// Output: [1 2 3]
}

func TestSort(t *testing.T) {
	list := []int{3, 1, 2}

	got := Reverse(list)
	expected := []int{1, 2, 3}
	if expected[0] != got[0] && expected[2] != got[2] {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func ExampleMap() {
	list := []int{3, 1, 2}

	fmt.Println(Map(list, func(item int) int {
		return item + 2
	}))
	//
	// Output: [5 3 4]
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

func ExampleFilter() {
	list := []int{3, 1, 2}

	fmt.Println(Filter(list, func(item int) bool {
		return item == 3
	}))
	//
	// Output: [3]
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

func ExampleFind() {
	list := []int{3, 1, 2}

	fmt.Println(Find(list, func(item int) bool {
		return item%3 == 0
	}))
	//
	// Output: 3 <nil>
}

func TestFind(t *testing.T) {
	list := []int{3, 1, 2}

	got, err := Find(list, func(e int) bool {
		return e%3 == 0
	})
	if err != nil {
		t.Fatalf("expected not to error")
	}

	expected := 3
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

	_, err = Find(list, func(e int) bool {
		return e == 100
	})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func ExampleAny() {
	list := []int{3, 1, 2}

	fmt.Println(Any(list, func(item int) bool {
		return item%3 == 0
	}))
	//
	// Output: true
}

func TestAny(t *testing.T) {
	list := []int{3, 1, 2}

	got := Any(list, func(e int) bool {
		return e%3 == 0
	})

	expected := true
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

	got = Any(list, func(e int) bool {
		return e == 100
	})

	expected = false
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func ExampleReduce() {
	list := []int{3, 1, 2}

	fmt.Println(Reduce(list, 0, func(cur, val int) int {
		return cur + val
	}))
	//
	// Output: 6
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

func ExampleReduceRight() {
	list := []int{2, 5, 10}

	fmt.Println(ReduceRight(list, 20, func(cur, val int) int {
		return cur - val
	}))
	//
	// Output: 3
}

func TestReduceRight(t *testing.T) {
	list := []int{2, 5, 10}

	got := ReduceRight(list, 20, func(cur, val int) int {
		return cur - val
	})
	expected := 3
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

	got = ReduceRight(list, 10, func(cur, val int) int {
		return cur - val
	})
	expected = -7
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

}
