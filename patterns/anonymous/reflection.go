package anonymous

import (
	"encoding/json"
	"reflect"
)

type TestStruct struct {
	Slice []json.RawMessage
}

func (t TestStruct) UnmarshalStruct(v interface{}) error {
	// get the value of the underlying slice
	slice := reflect.ValueOf(v).Elem()
	// make sure we have adequate capacity
	slice.Set(reflect.MakeSlice(slice.Type(), len(t.Slice), len(t.Slice)))

	for i, val := range t.Slice {
		err := json.Unmarshal(val, slice.Index(i).Addr().Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

type OtherType struct {
	Foo string
}
