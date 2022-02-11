package collection

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
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
