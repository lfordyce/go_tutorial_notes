package factory

import (
	"fmt"
	"testing"
)

func TestNewPersonFactory(t *testing.T) {
	newBaby := NewPersonFactory(5)
	person := newBaby("bruce")

	fmt.Println(person)

}
