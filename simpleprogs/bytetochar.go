package simpleprogs

import (
	"fmt"
	"unicode/utf8"
)

// func printchar(r rune) {
// 	fmt.Printf("%c is the char \n", r)
// }

func bytetorune(b []byte) rune {

	r, _ := utf8.DecodeRune(b)
	//fmt.Printf("The rune = %c \n", r)
	fmt.Printf("%c", r)
	return r
}
