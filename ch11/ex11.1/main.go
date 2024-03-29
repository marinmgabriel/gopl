// Exercise 11.1: Write tests for the charcount program in Section 4.3.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Charcount computes counts of Unicode characters.
func charcount(r io.Reader) (counts map[rune]int, utflen []int, invalid int) {
	counts = make(map[rune]int)         // counts of Unicode characters
	utflen = make([]int, utf8.UTFMax+1) // count of lengths of UTF-8 encodings
	invalid = 0                         // count of invalid UTF-8 characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return
}

func main() {
	counts, utflen, invalid := charcount(os.Stdin)
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
