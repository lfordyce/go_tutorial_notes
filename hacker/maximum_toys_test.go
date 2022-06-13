package hacker

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMaximumToys(t *testing.T) {
	input := []byte("7 50\n1 12 5 111 200 1000 10")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write(input); err != nil {
		t.Fatal(err)
	}

	stdin := os.Stdin
	// Restore stdin right after the test.
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	result := toyPriceSetup(maximumToys, t)
	t.Log(result)
}

func toyPriceSetup(fn func([]int32, int32) int32, t *testing.T) int32 {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	n := int32(nTemp)
	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	k := int32(kTemp)

	pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var prices []int32

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}
	return fn(prices, k)
}
