package hacker

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestTriangleOrNot(t *testing.T) {
	inputAlt := []byte("3\n7\n10\n7\n3\n2\n3\n4\n3\n2\n7\n4")
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

	result := testRunnerTriangels(triangleOrNot, t)
	t.Log(result)
}

func TestTriangleOrNot_case2(t *testing.T) {
	inputAlt := []byte("7\n10\n1\n9\n4\n5\n6\n10\n7\n5\n7\n3\n9\n4\n10\n10\n7\n5\n4\n8\n10\n1\n8\n8")
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

	result := testRunnerTriangels(triangleOrNot, t)
	t.Log(result)
}

func testRunnerTriangels(fn func(a []int32, b []int32, c []int32) []string, t *testing.T) []string {
	reader := bufio.NewReader(os.Stdin)

	aSize, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	var a []int32
	for i := 0; i < int(aSize); i++ {
		aItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		aItem := int32(aItemTemp)

		a = append(a, aItem)
	}

	bSize, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	var b []int32
	for i := 0; i < int(bSize); i++ {
		bItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		bItem := int32(bItemTemp)

		b = append(b, bItem)
	}

	cSize, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	var c []int32
	for i := 0; i < int(cSize); i++ {
		cItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		cItem := int32(cItemTemp)

		c = append(c, cItem)
	}
	return fn(a, b, c)
}
