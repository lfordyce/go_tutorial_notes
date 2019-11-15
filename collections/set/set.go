package set

type (
	Set struct {
		hash map[interface{}]nothing
	}
	nothing struct{}
)

func New(initial ...interface{}) *Set {
	s := &Set{map[interface{}]nothing{}}

	for _, v := range initial {
		s.Insert(v)
	}
	return s
}

// Find the difference of the two sets
func (s *Set) Difference(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = nothing{}
		}
	}
	return &Set{n}
}

// Call f for each item in the set
func (s *Set) Do(f func(interface{})) {
	for k := range s.hash {
		f(k)
	}
}

func (s *Set) Has(element interface{}) bool {
	_, exists := s.hash[element]
	return exists
}

func (s *Set) Insert(element interface{}) {
	s.hash[element] = nothing{}
}

func (s *Set) Intersection(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = nothing{}
		}
	}
	return &Set{n}
}

func (s *Set) Len() int {
	return len(s.hash)
}
