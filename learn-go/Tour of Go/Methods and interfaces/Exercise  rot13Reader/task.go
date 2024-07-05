package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return ((b - 'A' + 13) % 26) + 'A'
	} else if 'a' <= b && b <= 'z' {
		return ((b - 'a' + 13) % 26) + 'a'
	}

	// TODO: handle this
	return b
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	if err == nil {
		for i := range p {
			p[i] = rot13(p[i])
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
