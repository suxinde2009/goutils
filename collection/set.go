package collection

import "sync"

type Set struct {
	m   map[interface{}]struct{}
	len int
	sync.RWMutex
}

func NewSet(cap int64) *Set {
	temp := make(map[interface{}]struct{}, cap)
	return &Set{
		m: temp,
	}
}

func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.m[item] = struct{}{}
	s.len = len(s.m)
}

func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.len == 0 {
		return
	}

	delete(s.m, item)
	s.len = len(s.m)
}

func (s *Set) Has(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return s.len
}

func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.m = make(map[interface{}]struct{})
	s.len = 0
}

func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()

	list := make([]interface{}, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
