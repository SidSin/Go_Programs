package list

import "testing"

func TestListUnion01(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}
	list3 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)

	list3.AddatEnd(1)
	list3.AddatEnd(3)
	list3.AddatEnd(5)
	list3.AddatEnd(2)
	list3.AddatEnd(4)
	list3.AddatEnd(6)

	list1.ListUnion(&list2)

	if !ListsEqual(&list1, &list3) {
		t.Errorf("Both lists are same .Should be equal.")
	}
}

func TestListUnion02(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}
	list3 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)

	list3.AddatEnd(1)
	list3.AddatEnd(3)
	list3.AddatEnd(5)
	list3.AddatEnd(2)
	list3.AddatEnd(4)
	list3.AddatEnd(7)

	list1.ListUnion(&list2)

	if ListsEqual(&list1, &list3) {
		t.Errorf("Both lists are NOT same .Diff in last element.")
	}

	PrintLinkedList(&list3)
}

func TestFastListUnion01(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}
	list3 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)

	list3.AddatEnd(1)
	list3.AddatEnd(3)
	list3.AddatEnd(5)
	list3.AddatEnd(2)
	list3.AddatEnd(4)
	list3.AddatEnd(6)

	list1.FastListUnion(&list2)

	if !ListsEqual(&list1, &list3) {
		t.Errorf("Both lists are same .Should be equal.")
	}
}

func TestFastListUnion02(t *testing.T) {

	list1 := LinkedList{nil, nil}
	list2 := LinkedList{nil, nil}
	list3 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)

	list3.AddatEnd(1)
	list3.AddatEnd(3)
	list3.AddatEnd(5)
	list3.AddatEnd(2)
	list3.AddatEnd(4)
	list3.AddatEnd(7)

	list1.FastListUnion(&list2)

	if ListsEqual(&list1, &list3) {
		t.Errorf("Both lists are NOT same .Diff in last element.")
	}

	PrintLinkedList(&list3)
}
