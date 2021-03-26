package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(b byte) byte {
	if 'a' <= b && b <= 'z' {
		base := b - 'a'
		b = 'a' + (base+13)%26
	} else if 'A' <= b && b <= 'Z' {
		base := b - 'A'
		b = 'A' + (base+13)%26
	}
	return b
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
