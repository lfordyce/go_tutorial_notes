package concurrency

import "testing"

func TestProcess(t *testing.T) {
	// GIVEN
	input := make(chan string)
	defer close(input)

	done := make(chan bool)
	defer close(done)

	go func() {
		input <- "hello world"
		done <- true
	}()

	// WHEN
	output := Process(input)
	<-done // blocks until the input write routine is finished

	// THEN
	expected := "(hello world)"
	found := <-output // blocks until the output has contents

	if found != expected {
		t.Errorf("Expected %s, found %s", expected, found)
	}
}
