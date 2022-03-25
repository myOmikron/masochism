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
