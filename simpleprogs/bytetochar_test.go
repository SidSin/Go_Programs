package simpleprogs

import (
	"testing"
)

func TestBytetorune_01(t *testing.T) {

	b := []byte{97}
	r := bytetorune(b)

	if r != 'a' {
		t.Errorf("97 should map to a")
	}
}

func TestBytetorune_02(t *testing.T) {

	b := []byte{98}
	r := bytetorune(b)

	if r != 'b' {

		t.Errorf("98 should map to b")

	}
}

func TestBytetorune_03(t *testing.T) {

	b := []byte{'@'}
	r := bytetorune(b)

	if r != '@' {

		t.Errorf("0040 should map to @")

	}
}

func TestBytetorune_04(t *testing.T) {

	b := []byte{'*'}
	r := bytetorune(b)

	if r != '*' {

		t.Errorf("002A should map to *")

	}
}

func TestBytetorune_05(t *testing.T) {

	hexvalue := 0x61

	b := []byte{byte(hexvalue)}
	r := bytetorune(b)

	if r != 'a' {

		t.Errorf("0x61 should map to a")

	}
}

func TestBytetorune_06(t *testing.T) {

	hexvalue := 0x21 //!
	givenrune := '!'

	b := []byte{byte(hexvalue)}
	r := bytetorune(b)

	if r != givenrune {

		t.Errorf("%#x should map to %c", hexvalue, givenrune)

	}
}

func TestBytetorune_07(t *testing.T) {

	hexvalue := 0x3F //!
	givenrune := '?'

	b := []byte{byte(hexvalue)}
	r := bytetorune(b)

	if r != givenrune {

		t.Errorf("%#x should map to %c, found %c", hexvalue, givenrune, r)

	}
}

func TestBytetorune_08(t *testing.T) {

	hexvalue := 0x2E
	givenrune := '.'

	b := []byte{byte(hexvalue)}
	r := bytetorune(b)

	if r != givenrune {

		t.Errorf("%#x should map to %c, found %c", hexvalue, givenrune, r)

	}
}
