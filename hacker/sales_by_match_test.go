package hacker

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func _TestSalesByMatch(t *testing.T) {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func TestStdinPipe(t *testing.T) {
	input := []byte("9\n10 20 20 10 10 30 50 10 20")
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

	work := doWork(t)
	t.Log(work)
}

func doWork(t *testing.T) int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	//writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}
	return sockMerchant(n, ar)
	//fmt.Fprintf(writer, "%d\n", result)
	//writer.Flush()
}

// imagineMain is the main() function, because otherwise the Go playground would
// execute that instead of the tests which are the purpose of this file.
func imagineMain() {
	fmt.Println("Hello, foo")
	fmt.Fprintln(os.Stderr, "foo, bar")
}

// readAll wraps ioutil.ReadAll for asynchronous use in tests.
func readAll(t *testing.T, r io.Reader, c chan<- []byte) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	c <- data
}

func Test_mainWithSinglePipe(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	c := make(chan []byte)
	go readAll(t, r, c)

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	want := "Hello, foo\nfoo, bar\n"

	imagineMain()

	w.Close()

	if got := <-c; !bytes.Equal(got, []byte(want)) {
		t.Errorf("want=%q, got=%q", want, string(got))
	}
}

func Test_mainWithSeparatePipes(t *testing.T) {
	rStdout, wStdout, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rStderr, wStderr, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	cStdout := make(chan []byte)
	cStderr := make(chan []byte)
	go readAll(t, rStdout, cStdout)
	go readAll(t, rStderr, cStderr)

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = wStdout
	os.Stderr = wStderr

	wantOut := "Hello, foo\n"
	wantErr := "foo, bar\n"

	imagineMain()

	wStdout.Close()
	wStderr.Close()

	if gotOut := <-cStdout; !bytes.Equal(gotOut, []byte(wantOut)) {
		t.Errorf("want=%q, got=%q", wantOut, string(gotOut))
	}
	if gotErr := <-cStderr; string(gotErr) != wantErr {
		t.Errorf("want=%q, got=%q", wantErr, string(gotErr))
	}
}

func TestUserInputFromStdin(t *testing.T) {
	input := []byte("Alice")
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

	username, err := userInput()
	if err != nil {
		t.Fatalf("userInput: %v", err)
	}
	t.Log(username)
}

func userInput() (username string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		username = scanner.Text()
	}
	if err = scanner.Err(); err != nil {
		return "", err
	}
	return
}
