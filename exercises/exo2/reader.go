package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (sE spaceEraser) Read(p []byte) (n int, err error) {
	n, err = sE.r.Read(p)
	var final []byte
	for _, caractere := range p {
		if caractere != 32 { //si le caractère n'est pas un espace
			final = append(final, caractere)
		}
	}
	copy(p, final)
	return n, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
