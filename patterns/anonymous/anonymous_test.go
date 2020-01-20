package anonymous

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestAnonymousAbstractThing(t *testing.T) {
	app := App{}

	// we use an anonymous function which satisfies the interface
	// The trick here is to pass the function to our DoThingWith type
	app.DoThing(DoThingWithFunc(func() {
		fmt.Println("Hey interface, are you satisfied?")
	}))
}

func TestProtoThing(t *testing.T) {
	thing := newThing()

	fmt.Println(thing.Item())
	thing.SetItem(4.0)
	fmt.Println(thing.Item())
}

func TestInterface(t *testing.T) {
	var i I = &T2{"foo", T1{field1: "bar"}}
	i.M()
	fmt.Println(i.(*T2).field1)
	fmt.Println(i.(*T2).field2)
}

func TestBindFunc_Error(t *testing.T) {

	b := bindFunc(add)
	fmt.Println(b(5, 6))

	fmt.Println(b.Error())

	var times2 Multiply = multiply.Apply(2)
	fmt.Println("times 2:", times2(3, 4), "(expect 24)")

	apply := adding.Apply(3)
	fmt.Println("add 3 to 2:", apply(2))

	var addition Add = func(i int, k int) int {
		return i + k
	}

	///
	option := Third(3)

	options := applyOptions(addition, option)

	i := options.Apply(2)
	fmt.Println("Add with third option: 3 + 2 + 4: ", i(4))

	f := func(c rune) bool {
		return c == ':' || c == '/'
	}

	fieldsFunc := strings.FieldsFunc("controller://10.208.127.11", f)
	fmt.Println(fieldsFunc)
}

func TestAmount(t *testing.T) {
	//amount := currency.EUR.Amount()
	//formatter := currency.Symbol.Default(currency.EUR)
	//
	//value := formatter(9.0)
	//fmt.Println(value)

	parse, _ := url.Parse("halo://denprdpkg0001")

	i := EncoderURL.Apply(parse)("http")

	i2 := PackagerURL.Apply(parse)("http", "my-service-id")

	fmt.Println(i)
	fmt.Println(i2)
}
