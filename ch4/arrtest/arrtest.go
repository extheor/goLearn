package main

import "fmt"

func main() {
	q := [...]int{1, 2, 3}
	q2 := &q

	q2[0] = 66
	fmt.Println(q, *q2)
	fmt.Printf("q的地址=%p, q2的地址=%p", &q, q2)
}
