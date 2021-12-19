package list

import (
	"fmt"
	"testing"
)

func TestNodeCreate(t *testing.T) {

	n1 := Node{1, nil, -3}
	fmt.Println(n1)

	linkedlist1 := LinkedList{nil, nil}
	fmt.Println(linkedlist1)

	linkedlist1.AddatEnd(1)
	PrintLinkedList(&linkedlist1)

	linkedlist1.AddatEnd(2)
	PrintLinkedList(&linkedlist1)
}
