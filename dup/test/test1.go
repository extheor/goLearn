// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if filenames[input.Text()] == arg+" " {
			continue
		} else {
			filenames[input.Text()] += arg + " "
		}

	}
}
