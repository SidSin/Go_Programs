package subsetsum

import (
	"fmt"
	"sidsin/srcmodule/list"
)

//bruteforce
func ExactSubsetSum(L *list.LinkedList, target int) bool {

	fmt.Println("starting ExactSubsetSum")
	n := list.GetListLenght(L)
	currnode := L.GetListHead()

	fmt.Println("1")

	list0 := list.CreateEmptyLinkedList()
	list0.AddatEnd(0)
	listx := list0
	listz := list0

	fmt.Println("2")

	for i := 1; i <= n; i++ {

		fmt.Println("")
		fmt.Println("---- i=", i, " ----")
		fmt.Println("Printing listx")
		list.PrintLinkedList(listx)

		e := currnode.GetNodeValue()

		listx = calculatesums(e, target, listx)

		listz = listx
		currnode = currnode.GetNextNode()

	}

	foundexactsum := false

	if target == listz.GetListTailValue() {
		foundexactsum = true
	}

	return foundexactsum
}

func calculatesums(e, target int, listx *list.LinkedList) *list.LinkedList {

	fmt.Println("New e =", e)

	listy := listx.CopyList()
	listy.IncrementListElements(e)

	fmt.Println("Printing listy")
	list.PrintLinkedList(listy)

	listz := list.ListMergeSort(listx, listy)
	listz = list.RemoveDuplicateElements(listz)
	listz = list.RemoveElementsBiggerThan(target, listz)

	fmt.Println("Printing listz")
	list.PrintLinkedList(listz)

	return listz
}
