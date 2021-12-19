package subsetsum

import (
	"fmt"
	"sidsin/srcmodule/list"
	"testing"
)

func TestExactSubsetSum01(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list1.AddatEnd(3)
	list1.AddatEnd(4)
	list1.AddatEnd(3)
	list1.AddatEnd(1)

	fmt.Println("Printing given list.")
	list.PrintLinkedList(list1)

	if !ExactSubsetSum(list1, 6) {
		t.Errorf("Expected sum should be 6.")
	}

}

func TestExactSubsetSum02(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list1.AddatEnd(1)
	list1.AddatEnd(4)
	list1.AddatEnd(5)

	fmt.Println("Printing given list.")
	list.PrintLinkedList(list1)

	if ExactSubsetSum(list1, 60) {
		t.Errorf("Target 60 is not possible with {1,4,5}")
	}

}

func TestExactSubsetSum03(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list1.AddatEnd(1)
	list1.AddatEnd(4)
	list1.AddatEnd(5)

	fmt.Println("Printing given list.")
	list.PrintLinkedList(list1)

	if !ExactSubsetSum(list1, 10) {
		t.Errorf("Target 10 is possible with {1,4,5}")
	}

}

func TestExactSubsetSum04(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list1.AddatEnd(1)
	list1.AddatEnd(4)
	list1.AddatEnd(5)

	fmt.Println("Printing given list.")
	list.PrintLinkedList(list1)

	if !ExactSubsetSum(list1, 0) {
		t.Errorf("Target 0 is possible with {1,4,5}")
	}

}
