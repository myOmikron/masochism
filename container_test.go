package masochism_test

import (
	"github.com/myOmikron/masochism"
	"strings"
	"testing"
)

//TestForeach This function will test slices
func TestForeach(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b := make([]int, 0)
	masochism.Foreach(func(elem int) { b = append(b, elem) })(a)

	if len(a) != len(b) {
		t.Error("mismatch of slice len")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Error("elements order missmatch")
		}
	}
}

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b := masochism.Map(func(elem int) int { return elem * elem })(a)

	if len(a) != len(b) {
		t.Error("mismatch of slice len")
	}

	for i := 0; i < len(a); i++ {
		if b[i] != a[i]*a[i] {
			t.Error("function was not executed correctly")
		}
	}

}

func TestFilter(t *testing.T) {
	s := []string{"abc", "def", "fgh"}
	a := masochism.Filter(func(t string) bool { return strings.HasPrefix(t, "d") })(s)

	if len(a) != 1 {
		t.Error("length mismatch")
	}

	if a[0] != "def" {
		t.Error("Wrong element in slice")
	}
}

func TestAny(t *testing.T) {
	s := []string{"abc", "def", "fgh"}

	a := masochism.Any(func(t string) bool { return false })(s)

	if a {
		t.Error("Any should not be true")
	}

	a = masochism.Any(func(t string) bool { return "abc" == t })(s)

	if !a {
		t.Error("Any should be true")
	}
}

func TestEvery(t *testing.T) {
	s := []string{"a", "abc", "abcdef"}

	if !masochism.Every(func(t string) bool { return strings.HasPrefix(t, "a") })(s) {
		t.Error("Elements should start with \"a\"")
	}

	if masochism.Every(func(t string) bool { return strings.HasSuffix(t, "f") })(s) {
		t.Error("Elements does not end with \"f\"")
	}
}

func TestNone(t *testing.T) {
	s := []string{"a", "abc", "abcdef"}

	if !masochism.None(func(t string) bool { return strings.HasPrefix(t, "b") })(s) {
		t.Error("Elements does not start with \"b\"")
	}

	if masochism.None(func(t string) bool { return strings.HasSuffix(t, "f") })(s) {
		t.Error("There is an element with suffix \"f\"")
	}
}
