package main

import "fmt"

func main() {
	fmt.Println(*newInt())
}

// func newInt() *int {
// 	return new(int)
// }

func newInt() *int {
	var dummy int
	return &dummy
}
