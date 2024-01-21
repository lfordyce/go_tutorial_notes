package hacker

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestEncircular(t *testing.T) {
	//input := []byte("aba\n10")
	inputAlt := []byte("3\nG\nL\nRGRG")
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

	result := testRunnerStr(doesCircleExist, t)
	t.Log(result)
	assert.Equal(t, []string{"NO", "YES", "YES"}, result)
}

func TestEncircular_Case2(t *testing.T) {
	//input := []byte("aba\n10")
	inputAlt := []byte("1\nGRGL")
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

	result := testRunnerStr(doesCircleExist, t)
	t.Log(result)
	assert.Equal(t, []string{"NO"}, result)
}

func testRunnerStr(fn func([]string) []string, t *testing.T) []string {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	commandsCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	var commands []string

	for i := 0; i < int(commandsCount); i++ {
		commandsItem := readLine(reader)
		commands = append(commands, commandsItem)
	}

	return fn(commands)
}
