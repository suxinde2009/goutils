package collection

type IntegerSet map[int64]struct{}

func NewIntegerSet() IntegerSet {
	return make(map[int64]struct{})
}

func (set IntegerSet) Add(v int64) {
	if _, ok := set[v]; ok {
		return
	}
	set[v] = struct{}{}
}

func (set IntegerSet) Contains(v int64) bool {
	if _, ok := set[v]; ok {
		return true
	}
	return false
}

func (set IntegerSet) Remove(v int64) {
	if _, ok := set[v]; !ok {
		return
	}
	delete(set, v)
}

func (set IntegerSet) Clone() IntegerSet {
	n := NewIntegerSet()
	for k, v := range set {
		n[k] = v
	}
	return n
}
