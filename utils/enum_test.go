package utils

import "testing"

func TestWeekday_String(t *testing.T) {

	got := Thursday.String()
	want := "Thursday"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWeekday_Weekend(t *testing.T) {
	got := Saturday.Weekend()
	if got != true {
		t.Errorf("got %v, want %t", got, true)
	}
}
