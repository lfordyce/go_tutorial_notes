package maps

import (
	"fmt"
	"sort"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	someMeta := make(map[string]string)
	someMeta["Provider"] = "bloomberg"
	someMeta["Provider_ID"] = "bloomberg.com"
	//someMeta[AssetID] = "BLMG0000000000000000"

	otherMeta := make(map[string]string)
	otherMeta["Version_Major"] = "1"
	otherMeta["Version_Minor"] = "0"
	otherMeta["AssetID"] = "BLMG0000000000000001"

	maps := mergeMaps(someMeta, otherMeta)
	for k, v := range maps {
		fmt.Printf("Key: %s | Value: %s\n", k, v)
	}
}

func TestSortMap(t *testing.T) {
	m := map[string]int{"hello": 10, "foo": 20, "bar": 20}
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sort.Sort(sort.IntSlice(a))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
}

func TestOtherMapSort(t *testing.T) {
	mim := myIntMap{
		1:  "one",
		11: "eleven",
		3:  "three",
	}
	for _, k := range mim.sort() {
		fmt.Println(mim[k])
	}
	msm := myStringMap{
		"b": "two",
		"s": "twenty",
		"e": "five",
	}

	for _, k := range msm.sort() {
		fmt.Println(msm[k])
	}
}

func TestExtractStructFields(t *testing.T) {
	i := &Package{Provider: "BLOOMBERG"}

	var fn PassObj = passer
	if err := Init(i, fn); err != nil {
		t.Fatal(err)
	}
}

func passer(kv *map[string]interface{}) map[string]string {
	stringsMap := make(map[string]string)
	for k, v := range *kv {
		fmt.Printf("KEY: %s, Value: %s\n", k, v)
		switch value := v.(type) {
		case string:
			stringsMap[k] = value
		}
	}
	return stringsMap
}
