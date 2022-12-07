package main

import "fmt"

func main() {
	slice := []string{"foo", "bar"}
	slice2 := []string{"foo", "bar"}
	fmt.Println(&slice == &slice2) // false

	s := "hello"
	s2 := s[:2]
	fmt.Println(s2)
}
