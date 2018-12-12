package interview

import (
	"fmt"
	"testing"
)

func TestGameOfThrones(t *testing.T) {

	test := "aabbaaa"
	thrones := gameOfThrones(test)
	fmt.Println(thrones)
}
