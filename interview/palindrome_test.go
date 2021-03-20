package interview

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGameOfThrones(t *testing.T) {

	test := "aabbaaa"
	thrones := gameOfThrones(test)
	fmt.Println(thrones)
}

type fizzbuzz int

func (x fizzbuzz) String() string {
	result := ""
	if x%3 == 0 {
		result += "Fizz"
	}
	if x%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(int(x))
	}
	return result
}

func TestFizzBuzz(t *testing.T) {
	for x := fizzbuzz(1); x <= 100; x++ {
		fmt.Println(x)
	}
}
