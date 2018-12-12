package interview

import (
	"fmt"
	"github.com/lfordyce/generalNotes/collections/set"
	"strings"
)

//var exists = struct {}{}
//
//type set struct {
//	m map[string]struct{}
//}
//
//func NewSet() *set {
//	s := &set{}
//	s.m = make(map[string]struct{})
//	return s
//}
//
//func (s *set) Add(value string) {
//	s.m[value] = exists
//}
//
//func (s *set) Remove(value string) {
//	delete(s.m, value)
//}
//
//func (s *set) Contains(value string) bool {
//	_, c := s.m[value]
//	return c
//}

// Complete the twoStrings function below.
func TwoStrings(s1 string, s2 string) string {

	//set1 := NewSet()
	//set2 := NewSet()
	//
	//for _, elem := range s1 {
	//	set1.Add(string(elem))
	//}
	//
	//for _, elem := range s2 {
	//	set2.Add(string(elem))
	//}

	charS1 := strings.Split(s1, "")
	charS2 := strings.Split(s2, "")

	i3 := set.New("h", "e", "l", "l", "0")
	i4 := set.New("h", "e", "y")
	i5 := i3.Intersection(i4)
	fmt.Print(i5)

	characters1 := make([]interface{}, len(charS1))
	characters2 := make([]interface{}, len(charS2))

	for i, s := range charS1 {
		characters1[i] = s
	}

	for i , s := range charS2 {
		characters2[i] = s
	}

	i := set.New(characters1...)
	i2 := set.New(characters2...)

	intersection := i.Intersection(i2)

	if intersection.Len() > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	inter := Intersection(strings.Split(s1, ""), strings.Split(s2, ""))
	if len(inter) > 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func Intersection(s1, s2 []string) (inter []string) {

	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	// Remove dups from slice
	inter = removeDuplicates(inter)
	return
}

func removeDuplicates(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}