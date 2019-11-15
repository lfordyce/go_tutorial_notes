package interview

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWaterBottle(t *testing.T) {
	b := NewWaterBottle()

	fmt.Println("Type of bottle is, ", reflect.TypeOf(b))
	fmt.Println("Liters:", b.Liters())
}
