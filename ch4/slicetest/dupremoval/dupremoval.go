package main

import "fmt"

func main() {
	s := []string{"a", "b", "b", "b", "b", "c", "a", "a", "d", "d"}
	fmt.Println(dupremoval(s))
}

// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func dupremoval(s []string) []string {
	for i := 0; i < len(s)-1; {
		ar := s[i] == s[i+1] // 是否相邻重复
		if ar {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}
