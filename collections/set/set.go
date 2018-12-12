package set

type (
	Set struct {
		hash map[interface{}]nothing
	}
	nothing struct {}
)

func New(initial ...interface{}) *Set {
	s := &Set{map[interface{}]nothing{}}

	for _, v := range initial {
		s.Insert(v)
	}
	return s
}

// Find the difference of the two sets
func (this *Set) Difference(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range this.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = nothing{}
		}
	}
	return &Set{n}
}

// Call f for each item in the set
func (this *Set) Do(f func(interface{})) {
	for k := range this.hash {
		f(k)
	}
}

func (this *Set) Has(element interface{}) bool {
	_, exists := this.hash[element]
	return exists
}

func (this *Set) Insert(element interface{}) {
	this.hash[element] = nothing{}
}

func (this *Set) Intersection(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range this.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = nothing{}
		}
	}
	return &Set{n}
}

func (this *Set) Len() int {
	return len(this.hash)
}

