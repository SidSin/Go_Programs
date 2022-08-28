package substitutioncipher

import (
	"bufio"
	"fmt"
	"os"
	"srcmodule/simpleprogs"
	"testing"
)

func TestShiftOne_01(t *testing.T) {

	initializeunshiftedrunes()
	sr := GetShiftedRune('A', 1)

	fmt.Printf("Shifted rune = %c  %v", sr, "\n")

	if sr != 'B' {
		t.Errorf("A+1 should give B. ")
	}
}

func TestShiftOne_02(t *testing.T) {

	initializeunshiftedrunes()
	sr := GetShiftedRune('Z', 1)

	answer := '['

	fmt.Printf("Shifted rune = %c  %v", sr, "\n")

	if sr != answer {
		t.Errorf("Expected = %c , found %c", answer, sr)
	}
}

func TestShiftOne_03(t *testing.T) {

	initializeunshiftedrunes()
	r := single_space

	sr := GetShiftedRune(r, 1)

	fmt.Printf("Shifted rune = %c  %v", sr, "\n")

	if sr != r {
		t.Errorf("Expected = %c , found %c", r, sr)
	}
}

func TestShiftTwo_01(t *testing.T) {

	initializeunshiftedrunes()
	sr := GetShiftedRune('A', 2)

	answer := 'C'

	fmt.Printf("Shifted rune = %c  %v", sr, "\n")

	if sr != answer {
		t.Errorf("Expected = %c , found = %c", answer, sr)
	}
}

func TestEncryptFile_03(t *testing.T) {

	name := "TestFile_04"

	f := simpleprogs.Openfile("C:\\Users\\SIDDHSIN\\" + name + ".txt")
	ef, err := os.Create("C:\\Users\\SIDDHSIN\\Enr_" + name + ".txt")
	simpleprogs.Check(err)

	initializeunshiftedrunes()

	shift := 1

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		encryptedstr := encryptstring(line, shift)

		simpleprogs.WriteStringtoFile(*ef, encryptedstr)
	}

}

func encryptstring(line string, shift int) string {

	runecount := 0
	encryptedstr := ""

	for _, r := range line {
		r = GetShiftedRune(r, shift)

		encryptedstr = encryptedstr + string(r)
		runecount++
	}
	fmt.Println("runecount (03) = ", runecount)

	return encryptedstr
}
