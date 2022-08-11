package main

import (
	"fmt"
)

type Set struct {
	list map[int]struct{} //empty structs occupy 0 memory
}

func (s *Set) Has(v int) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v int) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v int) {
	delete(s.list, v)
}

func (s *Set) Clear() {
	s.list = make(map[int]struct{})
}

func (s *Set) Size() int {
	return len(s.list)
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[int]struct{})
	return s
}

//optional functionalities

//AddMulti Add multiple values in the set
func (s *Set) AddMulti(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

type FilterFunc func(v int) bool

// Filter returns a subset, that contains only the values that satisfies the given predicate P
func (s *Set) Filter(P FilterFunc) *Set {
	res := NewSet()
	for v := range s.list {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set) Union(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

func (s *Set) Intersect(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from s, that doesn't exists in s2 (param)
func (s *Set) Difference(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}


func main() {
  mySet := NewSet()
  fmt.Println("has 5", mySet.Has(5)) //false
  mySet.Add(5)
  fmt.Println("has 5", mySet.Has(5)) //true
  mySet.Remove(5)
  fmt.Println("has 5", mySet.Has(5)) //false

  mySet.AddMulti(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

  p := func(v int) bool { return v >= 5 }
  subSet := mySet.Filter(p)
  fmt.Println("all >= 5", subSet)
}
