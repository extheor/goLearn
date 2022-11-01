package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
	
}