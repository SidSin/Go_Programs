package list

import (
	"fmt"
	"testing"
)

func TestMergeList(t *testing.T) {

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

	list1.ListUnion(&list2)

	PrintLinkedList(&list1)

	fmt.Println("Node count of list1 = ", GetListLenght(&list1))

}
