package main

import (
	"bytes"
	"fmt"
)

// IntSet ...
type IntSet struct {
	words []uint64
}

// Add ...
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	fmt.Printf("x:%d, word:%d, bit:%d\n", x, word, bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String ...
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	fmt.Println("vim-go")
	var is IntSet
	is.Add(63)
	is.Add(64)
	is.Add(65)
	is.Add(128)
	is.Add(1024)
	fmt.Println(is.String())
}
