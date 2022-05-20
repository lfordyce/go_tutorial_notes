package hacker

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCountingValleys(t *testing.T) {
	input := []byte("8\nUDDDUDUU")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write(input); err != nil {
		t.Fatal(err)
	}
	w.Close()

	stdin := os.Stdin
	// Restore stdin right after the test.
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	work := setup(countingValleys, t)
	t.Log(work)
}

func setup(fn func(int32, string) int32, t *testing.T) int32 {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	steps := int32(stepsTemp)

	path := readLine(reader)
	return fn(steps, path)
}