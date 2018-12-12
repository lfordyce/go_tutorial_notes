package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()
	if s.Len() != 0  {
		t.Errorf("Failed, invalid stack length.")
	}
	s.Destroy()
	return
}

func TestStackPushLength(t *testing.T) {
	s := New()
	s.Push(14)
	s.Push(42)
	s.Push("testing")
	s.Push([]byte("Viper"))
	stkLen := s.Len()
	if stkLen != 4 {
		t.Errorf("Failed, invalid stack length, got %d expected 4", len)
	}
	s.Destroy()
	return
}

func TestSafeStack(t *testing.T) {
	S := New()

	for i := 0; i < 10; i++ {
		go func(j int) {
			S.Push(j)
		}(i)
	}

	t.Logf("%d elements", S.Len())

	next := S.Pop()

	for next != nil {
		t.Logf("%d\n", next.(int))
		next = S.Pop()
	}
}
