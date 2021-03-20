package utils

import (
	"fmt"
	"github.com/dyninc/qstring"
	"net/url"
	"testing"
)

func TestUnmarshaller(t *testing.T) {
	testIO := []struct {
		inp      url.Values
		expected interface{}
	}{
		{url.Values{"names": []string{"foo", "bar"}}, nil},
		{make(url.Values), errNoNames},
	}

	s := &MarshalInterfaceTest{Names: []string{}}
	for _, test := range testIO {
		err := qstring.Unmarshal(test.inp, s)
		if err != test.expected {
			t.Errorf("Expected Unmarshaller to return %s, but got %s instead", test.expected, err)
		}
	}
}

func TestUmarshallWithDifferentParamNames(t *testing.T) {
	values := url.Values{
		"r": []string{"l3-000000-dir-01"},
		"t": []string{"DTV"},
		"p": []string{"180"},
	}
	var data fairPlayRequest
	if err := qstring.Unmarshal(values, &data); err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
}
