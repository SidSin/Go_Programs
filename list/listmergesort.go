package list

func ListMergeSort(L1, L2 *LinkedList) *LinkedList {

	l1length := GetListLenght(L1)
	l2length := GetListLenght(L2)

	//fmt.Println("L1 Len = ", l1length, " , L2 Len = ", l2length)

	mergedlist := LinkedList{nil, nil}

	currnode1 := L1.head
	currnode2 := L2.head

	if l1length == 0 && l2length == 0 {

	} else if l1length == 0 && l2length > 0 {
		mergedlist = appendlist2(currnode2, mergedlist)
	} else if l1length > 0 && l2length == 0 {
		mergedlist = appendlist1(currnode1, mergedlist)
	} else if l1length > 0 && l2length > 0 {

		x := 1
		y := 1

		for x <= l1length && y <= l2length {
			currnode1, currnode2, x, y, mergedlist = twowaymerge(currnode1, currnode2, mergedlist, x, y)
		}

		addremainingelements(x, l1length, currnode2, mergedlist, y, l2length, currnode1)
	}

	return &mergedlist

}

func twowaymerge(currnode1 *Node, currnode2 *Node, mergedlist LinkedList, x int, y int) (*Node, *Node, int, int, LinkedList) {

	if currnode1.value <= currnode2.value {
		mergedlist.AddatEnd(currnode1.value)
		currnode1 = currnode1.next
		x++
		//fmt.Println("New x = ", x)
	} else {
		mergedlist.AddatEnd(currnode2.value)
		currnode2 = currnode2.next
		y++
		//fmt.Println("New y = ", y)
	}
	return currnode1, currnode2, x, y, mergedlist
}

func addremainingelements(x int, l1length int, currnode2 *Node, mergedlist LinkedList, y int, l2length int, currnode1 *Node) {
	if x > l1length {
		appendlist2(currnode2, mergedlist)
	}

	if y > l2length {
		appendlist1(currnode1, mergedlist)
	}
}

func appendlist1(currnode1 *Node, mergedlist LinkedList) LinkedList {
	for currnode1 != nil {
		mergedlist.AddatEnd(currnode1.value)
		currnode1 = currnode1.next
	}

	return mergedlist
}

func appendlist2(currnode2 *Node, mergedlist LinkedList) LinkedList {
	for currnode2 != nil {
		mergedlist.AddatEnd(currnode2.value)
		//fmt.Println("Added ", currnode2.value, "to mergedlist")
		currnode2 = currnode2.next
	}

	return mergedlist
}
