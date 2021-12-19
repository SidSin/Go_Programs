package subsetsum

import (
	"fmt"
	"sidsin/srcmodule/list"
	"testing"
)

func TestApproxSubSetSum01(t *testing.T) {
	list1 := list.CreateEmptyLinkedList()
	list1.AddatEnd(104)
	list1.AddatEnd(102)
	list1.AddatEnd(201)
	list1.AddatEnd(101)

	approxval := ApproxSubSetSum(list1, 308, 0.4)
	fmt.Println("approxval = ", approxval)

	if approxval != 302 {
		t.Errorf("Expected 302, found something else.")
	}
}

func TestApproxSubSetSum02(t *testing.T) {
	list1 := list.CreateEmptyLinkedList()
	list1.AddatEnd(2)
	list1.AddatEnd(4)
	list1.AddatEnd(6)
	list1.AddatEnd(8)
	list1.AddatEnd(10)
	list1.AddatEnd(12)
	list1.AddatEnd(14)

	approxval := ApproxSubSetSum(list1, 30, 0.4)
	fmt.Println("approxval = ", approxval)

	if approxval != 30 {
		t.Errorf("Expected 30, found something else.")
	}
}
