package json_utils

import (
	"fmt"
	"testing"
)

func TestAnalyseText(t *testing.T) {
	rt := AnalyseText("bloomberg.com")
	if rt.vowels != 4 {
		t.Fatalf("expected: 3, actual: %d", rt.vowels)
	}
	fmt.Printf("%+v", rt)
}
