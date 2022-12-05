package main

import "fmt"

func main() {
	r := compareStr("123", "213")
	fmt.Println(r)
}

// 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func compareStr(s1 string, s2 string) bool {
	sb1 := []byte(s1)
	sb2 := []byte(s2)

	for i := range sb1 {
		if i < len(sb1)-1 && sb1[i] > sb1[i+1] {
			sb1[i], sb1[i+1] = sb1[i+1], sb1[i]
		}
	}

	for i := range sb2 {
		if i < len(sb2)-1 && sb2[i] > sb2[i+1] {
			sb2[i], sb2[i+1] = sb2[i+1], sb2[i]
		}
	}

	fmt.Println("sb1: ", string(sb1))
	fmt.Println("sb2: ", string(sb2))

	return string(sb1) == string(sb2)
}
