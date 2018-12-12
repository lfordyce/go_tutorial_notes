package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()

	for i := 0; i < 10; i++ {
		s.Push(i + 1)
	}

	if size := s.Size; size != 10 {
		t.Errorf("Wrong count, expected 10 but got %d", size)
	}
}