package subsequence

import (
	"sidsin/srcmodule/list"
	"testing"
)

func TestSubSeq01(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	list1.AddatEnd(4)
	list1.AddatEnd(2)
	list1.AddatEnd(3)
	list1.AddatEnd(2)
	list1.AddatEnd(4)
	list1.AddatEnd(1)
	list1.AddatEnd(2)
	list1.AddatEnd(3)
	list1.AddatEnd(1)

	list2.AddatEnd(1)
	list2.AddatEnd(1)

	if !IsSubsequence(list1, list2) {
		t.Errorf("{1,1} is a subseq of {4,2,3,2,4,1,2,3,1}")
	}
}

func TestSubSeq02(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	list1.AddatEnd(4)
	list1.AddatEnd(2)
	list1.AddatEnd(3)
	list1.AddatEnd(2)
	list1.AddatEnd(4)
	list1.AddatEnd(1)
	list1.AddatEnd(2)
	list1.AddatEnd(3)
	list1.AddatEnd(1)

	list2.AddatEnd(24)
	list2.AddatEnd(4)
	list2.AddatEnd(1)

	if IsSubsequence(list1, list2) {
		t.Errorf("{24,4,1} is NOT a subseq of {4,2,3,2,4,1,2,3,1}")
	}
}

func TestSubSeq03(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	list1.AddatEnd(1)
	list1.AddatEnd(2)
	list1.AddatEnd(3)
	list1.AddatEnd(2)
	list1.AddatEnd(4)
	list1.AddatEnd(1)
	list1.AddatEnd(2)

	list2.AddatEnd(1)
	list2.AddatEnd(1)
	list2.AddatEnd(1)

	if IsSubsequence(list1, list2) {
		t.Errorf("{1,1,1} is NOT a subseq of {1,2,3,2,4,1,2}")
	}
}

func TestSubSeq04(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	list1.AddatEnd(1)
	list1.AddatEnd(2)
	list1.AddatEnd(3)
	list1.AddatEnd(2)
	list1.AddatEnd(4)
	list1.AddatEnd(1)
	list1.AddatEnd(2)

	list2.AddatEnd(1)
	list2.AddatEnd(1)

	if !IsSubsequence(list1, list2) {
		t.Errorf("{1,1} is a subseq of {1,2,3,2,4,1,2}")
	}
}

func TestSubSeq05(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	list1.AddatEnd(1) //A
	list1.AddatEnd(2) //B
	list1.AddatEnd(3) //C
	list1.AddatEnd(2) //B
	list1.AddatEnd(4) //D
	list1.AddatEnd(1) //A
	list1.AddatEnd(2) //B

	list2.AddatEnd(2) //B
	list2.AddatEnd(3) //C
	list2.AddatEnd(2) //B
	list2.AddatEnd(1) //A

	if !IsSubsequence(list1, list2) {
		t.Errorf("{2,3,2,1} is a subseq of {1,2,3,2,4,1,2}")
	}
}

func TestSubSeq06(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	list1.AddatEnd(1) //A
	list1.AddatEnd(2) //B
	list1.AddatEnd(3) //C
	list1.AddatEnd(2) //B
	list1.AddatEnd(4) //D
	list1.AddatEnd(1) //A
	list1.AddatEnd(2) //B

	//list2 is nil

	if IsSubsequence(list1, list2) {
		t.Errorf("{0} is NOT a subseq of {1,2,3,2,4,1,2}")
	}
}

func TestSubSeq07(t *testing.T) {

	list1 := list.CreateEmptyLinkedList()
	list2 := list.CreateEmptyLinkedList()

	//list1 is nil

	list2.AddatEnd(1) //A
	list2.AddatEnd(2) //B
	list2.AddatEnd(3) //C
	list2.AddatEnd(2) //B

	if IsSubsequence(list1, list2) {
		t.Errorf("{1,2,3,2} is NOT a subseq of {0}")
	}
}
