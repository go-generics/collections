package collections_test

import (
	"testing"

	"github.com/go-generics/collections"
)

func TestSet_String(t *testing.T) {
	s2 := collections.NewSet(0, 1)
	s2_str := s2.String()

	if s2_str != "[0 1]" && s2_str != "[1 0]" {
		t.Fatalf("failed format %s", s2_str)
	}
}
