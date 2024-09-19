package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
)

//go:embed input.txt
var content embed.FS

func main() {
	f, err := content.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var answer int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil && err != io.EOF {
			panic(err)
		}
		line := scanner.Text()
		var v int
		for i := 0; i < len(line); i++ {
			if vv, ok := extractDigit(line[i:]); ok {
				v = vv * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if vv, ok := extractDigit(line[i:]); ok {
				v += vv
				break
			}
		}
		answer += v
	}
	fmt.Println("Answer: ", answer)
}

func extractDigit(line string) (int, bool) {
	if line[0] >= '0' && line[0] <= '9' {
		return int(line[0] - '0'), true
	}

	return 0, false
}
