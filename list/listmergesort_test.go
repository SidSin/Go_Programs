package list

import (
	"fmt"
	"testing"
)

func TestListMergeSort05(t *testing.T) {

	list1 := LinkedList{nil, nil}

	list2 := LinkedList{nil, nil}

	//PrintLinkedList(&list2)

	fmt.Println("starting LMS")

	list3 := ListMergeSort(&list1, &list2)

	fmt.Println("ending LMS")

	//PrintLinkedList(list3)

}

func TestListMergeSort04(t *testing.T) {

	list1 := LinkedList{nil, nil}

	list2 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)
	list1.AddatEnd(7)
	list1.AddatEnd(9)
	list1.AddatEnd(11)

	PrintLinkedList(&list2)

	fmt.Println("starting LMS")

	list3 := ListMergeSort(&list1, &list2)

	fmt.Println("ending LMS")

	PrintLinkedList(list3)

}

func TestListMergeSort03(t *testing.T) {

	list1 := LinkedList{nil, nil}

	list2 := LinkedList{nil, nil}

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)
	list2.AddatEnd(8)
	list2.AddatEnd(10)
	list2.AddatEnd(12)

	PrintLinkedList(&list2)

	fmt.Println("starting LMS")

	list3 := ListMergeSort(&list1, &list2)

	fmt.Println("ending LMS")

	PrintLinkedList(list3)

}

func TestListMergeSort02(t *testing.T) {

	list1 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	PrintLinkedList(&list1)

	list2 := LinkedList{nil, nil}

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)
	list2.AddatEnd(8)
	list2.AddatEnd(10)
	list2.AddatEnd(12)

	PrintLinkedList(&list2)

	fmt.Println("starting LMS")

	list3 := ListMergeSort(&list1, &list2)

	fmt.Println("ending LMS")

	PrintLinkedList(list3)

}

func TestListMergeSort01(t *testing.T) {

	list1 := LinkedList{nil, nil}

	list1.AddatEnd(1)
	list1.AddatEnd(3)
	list1.AddatEnd(5)

	PrintLinkedList(&list1)

	list2 := LinkedList{nil, nil}

	list2.AddatEnd(2)
	list2.AddatEnd(4)
	list2.AddatEnd(6)

	PrintLinkedList(&list2)

	fmt.Println("starting LMS")

	list3 := ListMergeSort(&list1, &list2)

	fmt.Println("ending LMS")

	PrintLinkedList(list3)

}
