package sortingprgrms

func SliceAreEqual(sp1, sp2 *[]int) (bool, string) {

	if len(*sp1) != len(*sp2) {
		return false, "Length not equal."
	}

	sl2 := *sp2

	for i, v := range *sp1 {
		if v != sl2[i] {
			return false, "Elements dont match."
		}
	}

	return true, "ok"
}
