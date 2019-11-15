package json_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestOmitEmpty(t *testing.T) {
	d := Dog{
		Breed: "retriever",
	}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

func TestZeroNilDifference(t *testing.T) {
	n := 3
	d1 := Restaurant{
		NumberOfCustomers: &n,
	}

	b, _ := json.Marshal(d1)
	fmt.Println(string(b))

	d2 := Restaurant{}
	b2, _ := json.Marshal(d2)
	fmt.Println(string(b2))
}

func TestUnmarshalling(t *testing.T) {
	bytes := []byte(`{"Breed":"retriever","WeightKg":5}`)
	var d Dog
	if err := json.Unmarshal(bytes, &d); err != nil {
		t.Fatal(err)
	}
	fmt.Println(d)
}

func TestOmitEmptyMarshal(t *testing.T) {

	dogWithNilDimension := Dog{
		WeightKg: 5,
		Breed:    "retriever",
	}
	b, err := json.Marshal(dogWithNilDimension)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	w := 3
	h := 5

	dogWithPointerToZeroDim := Dog{
		WeightKg: 5,
		Breed:    "retriever",
		Size:     &dimension{},
	}
	marshal, err := json.Marshal(dogWithPointerToZeroDim)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(marshal))

	dogWithDim := Dog{
		WeightKg: 5,
		Breed:    "retriever",
		Size: &dimension{
			Height: h,
			Width:  w,
		},
	}
	bytes, err := json.Marshal(dogWithDim)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bytes))

	//check([]byte(`{""}`))
}

func TestUnitUnmarshal(t *testing.T) {
	check([]byte(`{"string":"this is a string field", "integer":1234}`))
	check([]byte(`{"string":"this is a string field"}`))
}

func check(js []byte) {
	var t Test
	if err := json.Unmarshal(js, &t); err != nil {
		log.Fatal(err)
	}

	newJs, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(newJs))
}

func TestCustomUnmarshal(t *testing.T) {
	claims := `{"sub": "1234565432", "name":"Batman", "admin": true}`
	var userJwt *UserJwt
	err := json.Unmarshal([]byte(claims), &userJwt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(userJwt.Id)
}
