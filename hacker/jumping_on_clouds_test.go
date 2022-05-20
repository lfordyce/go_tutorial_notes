package hacker

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	input := []byte("7\n0 0 1 0 0 1 0")
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

	result := workerSetup(jumpingOnClouds, t)
	t.Log(result)
}

func workerSetup(fn func([]int32) int32, t *testing.T) int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	n := int32(nTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}
	return fn(c)
}
