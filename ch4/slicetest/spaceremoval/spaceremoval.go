package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []byte{'a', 'b', 'b', ' ', ' ', ' ', 'd', ' ', 'e', '\t', '\t', 'f', 'f'}
	fmt.Printf("%c", spaceremoval(s))
}

// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func spaceremoval(s []byte) []byte {
	for i := 0; i < len(s)-1; {
		ar := s[i] == s[i+1] && unicode.IsSpace(rune(s[i])) // 是否为相邻重复的空格
		if ar {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}
