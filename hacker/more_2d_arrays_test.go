package hacker

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestArrayManipulator(t *testing.T) {
	input := []byte("5 3 \n1 2 100\n2 5 100\n3 4 100\n")
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

	result := testRunnerTwoDimensional(arrayManipulation, t)
	t.Log(result)
}

func testRunnerTwoDimensional(fn func(n int32, queries [][]int32) int64, t *testing.T) int64 {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}
	return fn(n, queries)
}
