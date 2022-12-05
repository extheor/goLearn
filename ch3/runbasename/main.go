package main

import (
	"fmt"
	"golearn/ch3/basename"
)

func main() {
	// fmt.Println(basename.Basename("a/b/c.go")) // "c"
	// fmt.Println(basename.Basename("c.d.go"))   // "c.d"
	// fmt.Println(basename.Basename("abc"))      // "abc"

	fmt.Println(basename.Basename2("a/b/c.go")) // "c"
	fmt.Println(basename.Basename2("c.d.go"))   // "c.d"
	fmt.Println(basename.Basename2("abc"))      // "abc"
}
