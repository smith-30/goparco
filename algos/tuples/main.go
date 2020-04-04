package main

func main() {

}

type Set struct {
	integerMap map[int]bool
}

func NewSet() *Set {
	s := &Set{
		integerMap: make(map[int]bool),
	}
	return s
}

func (s *Set) AddElement(v int) {
	s.integerMap[v] = true
}

func (s *Set) DeleteElement(v int) {
	delete(s.integerMap, v)
}

func (s *Set) ContainElement(v int) bool {
	_, ok := s.integerMap[v]
	return ok
}

func (s *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := NewSet()
	for v, _ := range s.integerMap {
		if anotherSet.ContainElement(v) {
			intersectSet.AddElement(v)
		}
	}
	return intersectSet
}

func (s *Set) Union(anotherSet *Set) *Set {
	unionSet := NewSet()
	for v, _ := range s.integerMap {
		unionSet.AddElement(v)
	}
	for v, _ := range anotherSet.integerMap {
		unionSet.AddElement(v)
	}
	return unionSet
}
