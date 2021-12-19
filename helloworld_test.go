package helloworld

import "testing"

func TestHello(t *testing.T) {

	want := "Hello World!"
	got := Hworld()

	if want != got {
		t.Errorf("Hworld() = %q, want %q", got, want)
	}

}
