package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(os.Args)
	fmt.Println(os.Args[0:])
	fmt.Println(s)
}