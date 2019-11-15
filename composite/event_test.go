package composite

import (
	"fmt"
	"testing"
)

func TestEvents_Event(t *testing.T) {
	var stats Stats
	stats.cnt = 33

	stats.Category("aa").cnt = 66
	stats.Category("aa").Event("bb").value = 99

	fmt.Println(stats)
	fmt.Println(stats.cnt, stats.Category("aa").Event("bb").value)
}

func TestRomanNumeral(t *testing.T) {
	fmt.Println(romanNumeralDict()(10))
	fmt.Println(romanNumeralDict()(100))

	dict := romanNumeralDict()
	fmt.Println(dict(400))
}
