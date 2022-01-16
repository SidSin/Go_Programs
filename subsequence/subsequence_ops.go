package subsequence

import "sidsin/srcmodule/list"

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
