package list

import "fmt"

/*

-- AddatFront      -- O(1)
-- AddatEnd        -- O(1)

-- RemoveFromFront -- O(1)
-- RemoveFromEnd   -- O(N)

*/
//Node of a singly linked list
type Node struct {
	value int
	next  *Node
	link  int
}

//LinkedList defines a singly linked list
type LinkedList struct {
	head *Node
	tail *Node
}

func (N *Node) GetNodeValue() int {
	return N.value
}

func (N *Node) GetNextNode() *Node {
	return N.next
}

func CreateEmptyLinkedList() *LinkedList {

	list := LinkedList{nil, nil}

	return &list
}

func (L *LinkedList) GetListHead() *Node {
	return L.head
}

func (L *LinkedList) GetListTailValue() int {
	return L.tail.value
}

//IsLinkedListEmpty returns true if the given linked list is empty
func (L *LinkedList) IsLinkedListEmpty() bool {

	isempty := false

	//condition for emply list is both head and tail are nil
	if L.head == nil && L.tail == nil {
		isempty = true
	}

	return isempty
}

//PrintLinkedList prints the values in the linked list
func PrintLinkedList(L *LinkedList) {

	if L.IsLinkedListEmpty() {
		fmt.Println("The given Linked List has no elements")
	} else {
		currentnode := L.head
		for currentnode != nil {
			fmt.Print(currentnode.value, " --> ")
			currentnode = currentnode.next
		}

		if currentnode == nil {
			fmt.Print(" XXX ")
		}
		fmt.Println("")

	}
}

//AddatEnd adds a given value to the end of the list
func (L *LinkedList) AddatEnd(Value int) {

	//setup newnode
	newnode := Node{}
	newnode.value = Value

	if L.IsLinkedListEmpty() {
		//set both head and tail
		L.head = &newnode
	} else {
		L.tail.next = &newnode
	}

	//move the tail to point to newnode
	L.tail = &newnode

}

//RemoveFromFront removes the node pointed by the head
func (L *LinkedList) RemoveFromFront() {

	if !(L.IsLinkedListEmpty()) {

		currentnode := L.head
		//move head forward
		L.head = L.head.next
		currentnode.next = nil

		//fmt.Println("Removed node with value = ", currentnode.value)

		if L.head == nil {
			L.tail = nil
		}
	}
}

//adds elements of list2 at the end of list1
//iterative method
func (L1 *LinkedList) ListUnion(L2 *LinkedList) {

	currnode := L2.head

	for currnode != nil {
		L1.AddatEnd(currnode.value)
		currnode = currnode.next
	}
}

//adds elements of list2 at the end of list1
//iterative method
func (L1 *LinkedList) FastListUnion(L2 *LinkedList) {

	currnode := L2.head

	for currnode != nil {
		L1.AddatEnd(currnode.value)
		currnode = currnode.next
	}
}

func GetListLenght(L2 *LinkedList) int {

	currnode := L2.head
	nodecount := 0

	for currnode != nil {
		nodecount++
		currnode = currnode.next
	}
	return nodecount
}

func ListsEqual(L1, L2 *LinkedList) bool {

	listsequaldebug := 1

	arelistsequal := true

	listlen1 := GetListLenght(L1)
	listlen2 := GetListLenght(L2)

	if listsequaldebug > 0 {
		fmt.Println("    List1 Len = ", listlen1, " , List2 Len = ", listlen2)
	}

	if listlen1 != listlen2 {
		arelistsequal = false
		return arelistsequal
	}

	//here lists are of same lenght
	// can be zero also
	arelistsequal = comparelists(L1, L2, listsequaldebug, listlen1)

	return arelistsequal
}

func comparelists(L1 *LinkedList, L2 *LinkedList, listsequaldebug int, listlen1 int) bool {

	currnode1 := L1.head
	currnode2 := L2.head
	arelistsequal := true

	if currnode1 == nil && currnode2 == nil {

		if listsequaldebug > 0 {
			fmt.Println("    Both lists are empty, so are equal.")
		}

	} else {

		if listsequaldebug > 0 {
			fmt.Println("    Both lists have same length.")
		}

		arelistsequal = comparelistvalues(listlen1, currnode1, currnode2, listsequaldebug)

	}
	return arelistsequal
}

func comparelistvalues(listlen1 int, currnode1, currnode2 *Node, listsequaldebug int) bool {

	if listsequaldebug > 0 {
		fmt.Println("    starting comparelistvalues, listlen1 = ", listlen1)
		fmt.Println("    List1 node1 value = ", currnode1.value)
		fmt.Println("    List2 node1 value = ", currnode2.value)
	}

	arelistsequal := true
	x := 1
	for x <= listlen1 {
		if currnode1.value == currnode2.value {
			x++
			currnode1 = currnode1.next
			currnode2 = currnode2.next
		} else {
			arelistsequal = false
			break

		}
	}
	return arelistsequal
}

func (L *LinkedList) IncrementListElements(n int) *LinkedList {

	currnode := L.head

	for currnode != nil {
		currnode.value = currnode.value + n
		currnode = currnode.next
	}

	return L
}

func (L *LinkedList) BoundedILE(n, bound int) *LinkedList {

	currnode := L.head

	for currnode != nil {
		if currnode.value+n <= bound {
			currnode.value = currnode.value + n
		}
		currnode = currnode.next

	}

	return L
}

//makes a new list which is an exact copy of given list
func (L *LinkedList) CopyList() *LinkedList {

	newlist := LinkedList{nil, nil}
	currnode := L.head

	for currnode != nil {
		newlist.AddatEnd(currnode.value)
		currnode = currnode.next
	}

	return &newlist
}

//assumes that the list L is sorted
//creates a new list with unique values from L
//and returns the new list
func RemoveDuplicateElements(L *LinkedList) *LinkedList {

	newlist := LinkedList{nil, nil}
	currnode := L.head
	prevvalue := -1

	for currnode != nil {

		if prevvalue != currnode.value {
			newlist.AddatEnd(currnode.value)
		}
		prevvalue = currnode.value
		currnode = currnode.next
	}

	return &newlist
}

func RemoveElementsBiggerThan(b int, L *LinkedList) *LinkedList {
	newlist := LinkedList{nil, nil}
	currnode := L.head

	for currnode != nil {
		if currnode.value <= b {
			newlist.AddatEnd(currnode.value)
		}
		currnode = currnode.next
	}

	return &newlist
}

func TrimList(L *LinkedList, delta float32) *LinkedList {
	newlist := LinkedList{nil, nil}
	currnode := L.head
	last := currnode.value
	currnode = currnode.next
	newlist.AddatEnd(last)

	//start from 2nd element
	for currnode != nil {

		newelement := (float32(last) * (1 + delta))
		//fmt.Println("newelement = ", newelement)
		newintelement := int(newelement)

		//fmt.Println("currnode.value = ", currnode.value, " , newelement = ", newelement)

		if currnode.value > newintelement {
			newlist.AddatEnd(currnode.value)
			last = currnode.value
		}
		currnode = currnode.next
	}

	return &newlist
}
