package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()

		if r == '\n' || r == '\r' {
			break
		}
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsSpace(r) {
			counts["space"]++
		} else if unicode.IsPunct(r) {
			counts["punct"]++
		} else if unicode.IsNumber(r) {
			counts["number"]++
		} else if unicode.IsSymbol(r) {
			counts["symbol"]++
		} else if unicode.IsMark(r) {
			counts["mark"]++
		} else if unicode.IsDigit(r) {
			counts["diget"]++
		} else if unicode.IsPrint(r) {
			counts["print"]++
		} else if unicode.IsControl(r) {
			counts["control"]++
		} else if unicode.IsGraphic(r) {
			counts["graphic"]++
		} else {
			counts["invalid"]++
		}
	}
	fmt.Printf("counts: %v\n", counts)
}
