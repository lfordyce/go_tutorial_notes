package strategy

import (
	"fmt"
	"testing"
)

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

func TestStrategy(t *testing.T) {
	add := Operation{Addition{}}
	result := add.Operate(3, 5)

	if result != 8 {
		//t.Errorf("Addition operation failed, %s", result)
		t.Error("Addition operation failed")
	}
}

func TestSecretStrategyResult(t *testing.T) {
	strategic := New(4, 5)

	//multiply := func(a, b int) int {
	//	return a * b
	//}

	//var Adder Strategy = Strategy(adder)
	//var Multy Strategy = Strategy(multiy)

	//add := func(a, b int) int {
	//	return a + b
	//}

	strategic.SetStrategy(multiy)
	result := strategic.Result()

	strategic.SetStrategy(adder)
	i := strategic.Result()

	fmt.Println(i)
	fmt.Println(result)
}

func adder(x, y int) int {
	return x + y
}

func multiy(x, y int) int {
	return x * y
}

//type options struct {
//	third int
//}
//
//func (o *options) additions(a, b int) int {
//	o.third
//}

type Transformation interface {
	Transform()
}
