package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	m := NewTreeMap()
	cases := []struct {
		key   string
		value interface{}
	}{
		{"fish", 0},
		{"cat", 1},
		{"dog", 2},
		{"cats", 3},
		{"caterpillar", 4},
		{"cattle", 5},
		{"apple", 6},
		{"battle", 7},
	}

	for _, c := range cases {
		m.Insert(c.key, c.value)
	}

	for _, c := range cases {
		if !m.Exists(c.key) {
			t.Errorf("want map to contain %q", c.key)
		}
	}

	for _, prefix := range []string{
		"app",
		"cat",
		"bat",
	} {
		if !m.Prefix(prefix) {
			t.Errorf("want prefix %q", prefix)
		}
	}

	for _, c := range cases {
		if val := m.Get(c.key); val != c.value {
			t.Errorf("expected key %s to have value %v, got %v", c.key, c.value, val)
		}
	}
}

func TestByteTrie(t *testing.T) {
	t0 := New()
	if t0.Lookup([]byte{0}) != nil {
		t.Errorf("Trie contains data for missing key.")
	}

	t1 := t0.Insert([]byte{0}, []byte("test"))
	if t0.Lookup([]byte{0}) != nil {
		t.Errorf("Insert is not immutable")
	}

	if string(t1.Lookup([]byte{0})) != "test" {
		t.Errorf("Inserted value not found by look up.")
	}

	t2 := t1.Insert([]byte{1}, []byte("another test"))
	if string(t1.Lookup([]byte{0})) != "test" {
		t.Errorf("Insert is not immutable")
	}

	if t1.Lookup([]byte{1}) != nil {
		t.Errorf("Insert is not immutable")
	}

	if string(t2.Lookup([]byte{0})) != "test" {
		t.Errorf("Inserted value not found by look up.")
	}
	if string(t2.Lookup([]byte{1})) != "another test" {
		t.Errorf("Inserted value not found by look up.")
	}

	t3 := t2.Insert([]byte{0, 1}, []byte("a final test"))
	if string(t3.Lookup([]byte{0})) != "test" {
		t.Errorf("Inserted value not found by look up.")
	}
	if string(t3.Lookup([]byte{1})) != "another test" {
		t.Errorf("Inserted value not found by look up.")
	}
	if string(t3.Lookup([]byte{0, 1})) != "a final test" {
		t.Errorf("Inserted value not found by look up.")
	}

}
