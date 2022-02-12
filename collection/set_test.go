package collection

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet(5)

	s.Add(1)
	s.Add(2)
	s.Add(1)

	fmt.Println("list of all items in set: ", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("The set is cleared, now is empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 exists in the set")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items in set: ", s.List())
}

func TestIntegerSet(t *testing.T) {
	intSet := NewIntegerSet()

	intSet.Add(333)

	if intSet.Contains(333) {
		fmt.Println("Add function of IntegerSet successed.")
	} else {
		fmt.Println("Add function of IntegerSet failed.")
	}

	clonedIntSet := intSet.Clone()

	if clonedIntSet.Contains(333) {
		fmt.Println("Clone function of IntegerSet successed.")
	} else {
		fmt.Println("Clone function of IntegerSet failed.")
	}

	intSet.Remove(333)
	if intSet.Contains(333) {
		fmt.Println("Remove function of IntegerSet failed.")
	} else {
		fmt.Println("Remove function of IntegerSet successed.")
	}

}
