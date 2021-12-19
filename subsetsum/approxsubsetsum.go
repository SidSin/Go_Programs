package subsetsum

import (
	"fmt"
	"sidsin/srcmodule/list"
)

func ApproxSubSetSum(L *list.LinkedList, target int, e float32) int {

	list0 := list.CreateEmptyLinkedList()
	list0.AddatEnd(0)
	listx := list0
	listz := list0

	currnode := L.GetListHead()
	v := L.GetListHead().GetNodeValue()
	listlen := list.GetListLenght(L)

	for i := 1; i <= listlen; i++ {
		fmt.Println("")
		fmt.Println("---- i = ", i, " , v = ", v, " ----")

		fmt.Print("List x = ")
		list.PrintLinkedList(listx)

		listy := listx.CopyList()
		listy.IncrementListElements(v)
		fmt.Print("List y = ")
		list.PrintLinkedList(listy)

		listz = list.ListMergeSort(listx, listy)
		//list.PrintLinkedList(listz)
		listz = list.RemoveDuplicateElements(listz)
		listz = list.TrimList(listz, e/float32(2*listlen))
		//list.PrintLinkedList(listz)
		listz = list.RemoveElementsBiggerThan(target, listz)
		fmt.Print("List z = ")
		list.PrintLinkedList(listz)

		if currnode != nil {
			currnode = currnode.GetNextNode()
		}

		if currnode != nil {
			v = currnode.GetNodeValue()
		}

		listx = listz
	}

	return listz.GetListTailValue()
}
