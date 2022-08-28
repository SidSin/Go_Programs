package simpleprogs

import (
	"io"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Openfile(filename string) *os.File {

	f, err := os.Open(filename)
	Check(err)

	return f
}

func Readstringfromfile(f os.File) (string, error) {
	buf := make([]byte, 1)

	n, err := f.Read(buf)
	//At end of file, Read returns 0, io.EOF.

	if n == 0 {
		return string(""), io.EOF
	}

	return string(buf), err
}

func WriteStringtoFile(f os.File, str string) {
	_, err := f.WriteString(str)
	Check(err)
}

func Closefile(f os.File) {
	f.Close()
}
