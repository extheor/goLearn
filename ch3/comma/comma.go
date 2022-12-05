package main

import "fmt"

func main() {
	r := comma("12345678")
	fmt.Println(r)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
