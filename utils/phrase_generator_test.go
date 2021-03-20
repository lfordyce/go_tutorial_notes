package utils

import (
	"fmt"
	"github.com/sethvargo/go-diceware/diceware"
	"strings"
	"testing"
)

func TestPhraseGen(t *testing.T) {
	for i := 0; i < 20; i++ {
		list, err := diceware.Generate(3)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(strings.Join(list, "-"))
	}
}
