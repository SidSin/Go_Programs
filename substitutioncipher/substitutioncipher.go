package substitutioncipher

const (
	comma        = ','
	full_stop    = '.'
	single_space = ' '
	new_line     = '\u2424'
	new_line2    = '\n'
)

var unshifted_runes map[rune]bool

func initializeunshiftedrunes() {

	unshifted_runes = make(map[rune]bool)
	unshifted_runes[single_space] = true
	unshifted_runes[new_line] = true
	unshifted_runes[new_line2] = true
	unshifted_runes[comma] = true

}

func RetainOriginalRune(r rune) bool {

	return unshifted_runes[r]

}

func GetShiftedRune(r rune, shift int) rune {

	if RetainOriginalRune(r) {
		return r
	} else {
		return (r + rune(shift))
	}

}
