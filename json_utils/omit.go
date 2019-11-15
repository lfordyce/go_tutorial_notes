package json_utils

import (
	"strconv"
	"strings"
)

type dimension struct {
	Height int
	Width  int
}

type Dog struct {
	Breed    string
	WeightKg int
	// Now `size` is a pointer to a `dimension` instance
	Size *dimension `json:",omitempty"`
}

type Restaurant struct {
	NumberOfCustomers *int `json:",omitempty"`
}

type Test struct {
	FieldA *string `json:"string,omitempty"`
	FieldB *uint   `json:"integer,omitempty"`
}

type Sub uint

type UserJwt struct {
	Id    Sub `json:"sub"`
	Name  string
	Admin bool
}

func (s *Sub) UnmarshalJSON(b []byte) error {
	sub := strings.Replace(string(b), `"`, "", 2)
	v, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*s = Sub(uint(v))
	return nil
}
