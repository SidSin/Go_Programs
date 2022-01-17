package subsequence

import (
	"fmt"
	"sidsin/srcmodule/list"
)

func IsSubsequence(seq *list.LinkedList, subseq *list.LinkedList) bool {

	subseqfound := false

	e1 := seq.GetListHead()
	e2 := subseq.GetListHead()

	for (e1 != nil) && (e2 != nil) {
		if e1.GetNodeValue() == e2.GetNodeValue() {

			if e2 == subseq.GetListTail() {
				subseqfound = true
			}

			e1 = e1.GetNextNode()
			e2 = e2.GetNextNode()

			if e2 == nil {
				e1 = nil
			}

		} else {
			e1 = e1.GetNextNode()

			if e1 == nil {
				e2 = nil
			}
		}
	}

	return subseqfound
}

func LcsLenght(seq1, seq2 *list.LinkedList) {

	m := list.GetListLenght(seq1)
	n := list.GetListLenght(seq2)

	//subprob matrix
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}

	//for simplicity making b the same size as c
	//same indices can be used by b
	b := make([][]int, m+1)
	for i := range b {
		b[i] = make([]int, n+1)
	}

	matptr := &c
	bp := &b

	//PrintSubProbMatrix(matptr)

	calcmatrixelements(m, n, matptr, bp, seq1, seq2)

	PrintSubProbMatrix(matptr)

	fmt.Print("LCS = ")
	PrintLCS(bp, seq1, m, n)
	fmt.Println("")
}

func calcmatrixelements(m int, n int, p *[][]int, bp *[][]int, seq1, seq2 *list.LinkedList) {

	list1node := seq1.GetListHead()
	list2node := seq2.GetListHead()

	c := *p
	b := *bp

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			// fmt.Println("first")
			// fmt.Println(" i = ", i, ", j = ", j, " , N1 Value = ", list1node.GetNodeValue(), " , N2 Value = ", list2node.GetNodeValue())
			// fmt.Println("second")

			if list1node.GetNodeValue() == list2node.GetNodeValue() {
				c[i][j] = 1 + c[i-1][j-1]
				b[i][j] = 1
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = 2
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = 3
			}

			list2node = list2node.GetNextNode()
			//fmt.Println("last")
		}
		list1node = list1node.GetNextNode()
		//reset list2node to head for new value of i
		list2node = seq2.GetListHead()
	}
}

func PrintSubProbMatrix(p *[][]int) {
	fmt.Println("")

	c := *p

	for i := range c {
		for j := range c[i] {
			fmt.Print(c[i][j], " ")
		}
		fmt.Println("")
	}
}

func PrintLCS(bp *[][]int, seq1 *list.LinkedList, i, j int) {

	b := *bp

	if i == 0 || j == 0 {
		return
	} else if b[i][j] == 1 {
		PrintLCS(bp, seq1, i-1, j-1)
		fmt.Print(list.PrintNthNodeElement(seq1, i))
	} else if b[i][j] == 2 {
		PrintLCS(bp, seq1, i-1, j)
	} else {
		PrintLCS(bp, seq1, i, j-1)
	}

}
