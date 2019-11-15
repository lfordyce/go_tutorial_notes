package composite

import "testing"

var indata = `[
  {
    "ID": 1,
    "Name": "Root",
    "ParentID": 0,
    "Path": "Root"
  },
  {
    "ID": 2,
    "Name": "Ball",
    "ParentID": 1,
    "Path": "Root/Ball"
  },
  {
    "ID": 3,
    "Name": "Foot",
    "ParentID": 2,
    "Depth": 2,
    "Path": "Root/Ball/Foot"
  },
  {
    "ID": 4,
    "Name": "Soccer",
    "ParentID": 3,
    "Depth": 2,
    "Path": "Root/Ball/Soccer"
  }
]`

func TestMakeTree(t *testing.T) {
	if err := MakeTree(indata); err != nil {
		t.Fatal(err)
	}
}
