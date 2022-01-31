package simpleprogs

import "fmt"

const (
	fone uint = 1 << iota
	ftwo
	fthree
	ffour
	ffive
	fsix
)

const (
	monday uint = iota
	tuesday
	wednesday
)

func PrintConst() {
	fmt.Println("fone   = ", fone)
	fmt.Println("ftwo   = ", ftwo)
	fmt.Println("fthree = ", fthree)
	fmt.Println("ffour  = ", ffour)
	fmt.Println("ffive  = ", ffive)
	fmt.Println("fsix   = ", fsix)
}

func PrintWeekdayNumbers() {
	fmt.Println("monday    = ", monday)
	fmt.Println("tuesday   = ", tuesday)
	fmt.Println("wednesday = ", wednesday)
}
