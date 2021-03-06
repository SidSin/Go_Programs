package list

import (
	"fmt"
	"testing"
)

//1
func TestListsEqual01(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}

	if !ListsEqual(&list1, &list2) {
		t.Errorf("Both lists are null, should be equal.")
	}
}

//2
func TestListsEqual02(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}

	list2.AddatEnd(1)

	if ListsEqual(&list1, &list2) {
		t.Errorf("List1 has zero lenght, List2 has length = 1, should NOT be equal.")
	}
}

//3
func TestListsEqual03(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list2.AddatEnd(1)

	if !ListsEqual(&list1, &list2) {
		t.Errorf("Both lists have only one element 1.Should be equal.")
	}
}

//4
func TestListsEqual04(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list2.AddatEnd(2)

	if ListsEqual(&list1, &list2) {
		t.Errorf("Both lists have only one element and they are diff.Should NOT be equal.")
	}
}

//5
func TestListsEqual05(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	list2.AddatEnd(1)
	list2.AddatEnd(3)
	list2.AddatEnd(5)

	if !ListsEqual(&list1, &list2) {
		t.Errorf("Both lists have same elements.Should be equal.")
	}
}

//6
func TestListsEqual06(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	list2.AddatEnd(1)
	list2.AddatEnd(3)
	list2.AddatEnd(6)

	if ListsEqual(&list1, &list2) {
		t.Errorf("Both lists diff in last element.Should NOT be equal.")
	}
}

//7
func TestIncrementList01(t *testing.T) {

	list := LinkedList{nil, nil}
	list.AddatEnd(1)
	list.AddatEnd(2)
	list.AddatEnd(3)

	list.IncrementListElements(10)

	PrintLinkedList(&list)

}

//8
func TestIncrementList02(t *testing.T) {

	list := LinkedList{nil, nil}

	list.IncrementListElements(10)

	PrintLinkedList(&list)

}

//9
func TestGetListHead01(t *testing.T) {

	list := LinkedList{nil, nil}
	list.AddatEnd(23)

	headnode := list.GetListHead()
	fmt.Println("Value of headnode = ", headnode.value)
}

//10
func TestCreateEmptyLinkedList(t *testing.T) {

	list := CreateEmptyLinkedList()
	PrintLinkedList(list)

	list.AddatEnd(10)
	PrintLinkedList(list)

}

//11
func TestGetNodevalue01(t *testing.T) {

	node1 := Node{50, nil, 0}

	nodevalue := node1.GetNodeValue()
	fmt.Println("Value of node = ", nodevalue)
}

//12
func TestGetNextNode(t *testing.T) {

	node1 := Node{50, nil, 0}
	node2 := Node{60, nil, 0}

	node1.next = &node2

	node3 := node1.GetNextNode()
	fmt.Println("Value of node = ", node3.GetNodeValue())
}

//13
func TestRemoveDuplicateElements01(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(3)
	list1.AddatEnd(4)
	list1.AddatEnd(4)
	list1.AddatEnd(4)
	list1.AddatEnd(5)
	list1.AddatEnd(6)
	list1.AddatEnd(7)
	list1.AddatEnd(7)
	list1.AddatEnd(7)
	list1.AddatEnd(8)
	list1.AddatEnd(8)
	list1.AddatEnd(10)
	list1.AddatEnd(11)

	list2 := RemoveDuplicateElements(&list1)

	PrintLinkedList(list2)

}

//14
func TestRemoveDuplicateElements02(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(4)
	list1.AddatEnd(5)
	list1.AddatEnd(6)
	list1.AddatEnd(7)
	list1.AddatEnd(8)
	list1.AddatEnd(10)
	list1.AddatEnd(11)

	list2 := RemoveDuplicateElements(&list1)

	PrintLinkedList(list2)

}

//15
func TestTrimList01(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list1.AddatEnd(10)
	list1.AddatEnd(11)
	list1.AddatEnd(12)
	list1.AddatEnd(15)
	list1.AddatEnd(20)
	list1.AddatEnd(21)
	list1.AddatEnd(22)
	list1.AddatEnd(23)
	list1.AddatEnd(24)
	list1.AddatEnd(29)

	trimmedlistptr := TrimList(&list1, 0.1)

	list2 := LinkedList{nil, nil}
	list2.AddatEnd(10)
	list2.AddatEnd(12)
	list2.AddatEnd(15)
	list2.AddatEnd(20)
	list2.AddatEnd(23)
	list2.AddatEnd(29)

	listptr2 := &list2

	if !ListsEqual(trimmedlistptr, listptr2) {
		t.Errorf("Trimmed list is not same as {10,12,15,20,23,29}.")
	}
}

//16
func TestTrimList02(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list1.AddatEnd(10)
	list1.AddatEnd(11)
	list1.AddatEnd(12)
	list1.AddatEnd(15)
	list1.AddatEnd(20)
	list1.AddatEnd(21)
	list1.AddatEnd(22)
	list1.AddatEnd(23)
	list1.AddatEnd(24)
	list1.AddatEnd(29)

	trimmedlistptr := TrimList(&list1, 0.1)

	list2 := LinkedList{nil, nil}
	list2.AddatEnd(10)
	list2.AddatEnd(12)
	list2.AddatEnd(15)
	list2.AddatEnd(20)
	list2.AddatEnd(23)

	listptr2 := &list2

	if ListsEqual(trimmedlistptr, listptr2) {
		t.Errorf("Trimmed list is not same as {10,12,15,20,23}.")
		PrintLinkedList(trimmedlistptr)
	}
}

func TestPrintNthNodeElement01(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list1.AddatEnd(10)
	list1.AddatEnd(12)
	list1.AddatEnd(15)
	list1.AddatEnd(20)
	list1.AddatEnd(23)

	if PrintNthNodeElement(&list1, 2) != 12 {
		t.Errorf("Second element should be 12")
	}
}

func TestPrintNthNodeElement02(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list1.AddatEnd(10)
	list1.AddatEnd(12)
	list1.AddatEnd(15)
	list1.AddatEnd(20)
	list1.AddatEnd(23)

	if PrintNthNodeElement(&list1, 20) != 12 {
		t.Errorf("Second element should be 12")
	}
}
