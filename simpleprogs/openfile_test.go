package simpleprogs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"srcmodule/substitutioncipher"
	"strings"
	"testing"
)

func TestOpenfile(t *testing.T) {
	f := Openfile("C:\\Users\\SIDDHSIN\\TestFile_01.txt")

	str, e := Readstringfromfile(*f)

	Check(e)

	fmt.Println("Text from the file = ", str)

	Closefile(*f)
}

func TestReadmultiplelines_01(t *testing.T) {

	var str string
	var readerr error

	f := Openfile("C:\\Users\\SIDDHSIN\\TestFile_01.txt")

	//for {
	str, readerr = Readstringfromfile(*f)
	fmt.Println("Text from the file = ", str)
	if readerr == io.EOF {
		fmt.Println("Error in reading file.")
	}

	str, readerr = Readstringfromfile(*f)
	fmt.Println("Text from the file = ", str)
	if readerr == io.EOF {
		fmt.Println("Error in reading file.")
	}

	str, readerr = Readstringfromfile(*f)
	fmt.Println("3rd Text from the file = ", str)
	if readerr == io.EOF {
		fmt.Println("Error in reading file.")
	}

	//}

	Closefile(*f)
}

func TestReadmultiplelines_02(t *testing.T) {

	var str string
	var readerr error
	shift := 1

	var encryptedstr string

	f := Openfile("C:\\Users\\SIDDHSIN\\TestFile_02.txt")
	ef, err := os.Create("C:\\Users\\SIDDHSIN\\Enr_TestFile_02.txt")
	Check(err)

	for {
		str, readerr = Readstringfromfile(*f)

		if readerr == io.EOF {
			fmt.Println("") // needed for better output when test is run
			break
		}
		fmt.Print(" ---->", str)
		fmt.Print("\n")

		//str = strings.ToValidUTF8(str, " ")

		//fmt.Println("-------------------------")

		runecount := 0

		for _, r := range str {

			r = substitutioncipher.GetShiftedRune(r, shift)

			encryptedstr = encryptedstr + string(r)
			runecount++

		}

		fmt.Println("runecount = ", runecount)
		encryptedstr = strings.ToValidUTF8(encryptedstr, "")

		WriteStringtoFile(*ef, encryptedstr)

	}

	Closefile(*f)
	Closefile(*ef)
	//fmt.Println("File Closed")
}

func TestReadmultiplelines_03(t *testing.T) {

	f := Openfile("C:\\Users\\SIDDHSIN\\TestFile_03.txt")
	ef, err := os.Create("C:\\Users\\SIDDHSIN\\Enr_TestFile_03.txt")
	Check(err)

	shift := 1

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		encryptedstr := encryptstring(line, shift)

		WriteStringtoFile(*ef, encryptedstr)
	}

}

func encryptstring(line string, shift int) string {

	runecount := 0
	encryptedstr := ""

	for _, r := range line {
		r = substitutioncipher.GetShiftedRune(r, shift)

		encryptedstr = encryptedstr + string(r)
		runecount++
	}
	fmt.Println("runecount (03) = ", runecount)

	return encryptedstr
}
