package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot12Reader struct {
	R io.Reader
}

func (r rot12Reader) Read(b []byte) (n int, err error) {
	tn, tr := r.R.Read(b)
	if tr == nil {
		for v := 0; v < tn; v++ {
			i := b[v]
			fmt.Printf("----%v,%q,%v", v, b[v], b[v])
			fmt.Println()

			if (i >= 'N' && i <= 'Z') || (i >= 'n' && i <= 'z') {
				b[v] = i - 13
			} else if (i >= 'A' && i < 'N') || (i >= 'a' && i < 'n') {
				b[v] = i + 13
			}
			fmt.Printf("~~~~%v,%q,%v", v, b[v], b[v])
			fmt.Println()
		}
	}
	return tn, tr
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot12Reader{s}
	io.Copy(os.Stdout, &r)
}
