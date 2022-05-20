package hacker

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestRepeatedString(t *testing.T) {
	//input := []byte("aba\n10")
	inputAlt := []byte("a\n1000000000000")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write(inputAlt); err != nil {
		t.Fatal(err)
	}
	w.Close()

	stdin := os.Stdin
	// Restore stdin right after the test.
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	result := testSetup(repeatedStringAlt, t)
	t.Log(result)

}

func testSetup(fn func(string,int64)int64, t *testing.T) int64 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	s := readLine(reader)
	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	return fn(s, n)
}