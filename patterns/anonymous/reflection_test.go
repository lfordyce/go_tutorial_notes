package anonymous

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

var data = []byte(`{
	"Slice": [
		{"Foo": "item one"},
		{"Foo": "item two"}
	]
}`)

func getTestData() *TestStruct {
	ts := &TestStruct{}
	err := json.Unmarshal(data, ts)
	if err != nil {
		log.Fatal(err)
	}
	return ts
}

func TestReflection(t *testing.T) {
	ts := getTestData()

	var others []OtherType
	err := ts.UnmarshalStruct(&others)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(others)
}
