package main

import "fmt"

func main() {
	// a := [...]int{1, 2, 3, 4, 5}
	// reverse(a[:])
	// fmt.Println(a)

	// s := []int{0, 1, 2, 3, 4, 5}
	// reverse(s[:2])
	// reverse(s[2:])
	// reverse(s)
	// fmt.Println(s)

	// s := [...]int{0, 1, 2, 3, 4, 5}
	// reverse(&s)
	// fmt.Println(s)

	s := []byte{0, 1, 2, 3, 4, 5}
	reverse_rune(s)
	fmt.Println(s)

	var a []int
	fmt.Println(a == nil)
}

// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }

// 重写reverse函数，使用数组指针代替slice。
// func reverse(s *[6]int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }

// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？(测试结果是不用分配额外的内存也可以)
func reverse_rune(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
