package main

import "fmt"

func main() {
	p := new(int)
	q := new(int)
	fmt.Println(p == q) // false

	// var a struct{}
	// var b struct{}
	// fmt.Println(a == b) // true
}
