package main

import "fmt"

func main() {
	// var runes []rune
	// for _, r := range "Hello, 世界" {
	// 	runes = append(runes, r)
	// }
	// fmt.Printf("%q\n", runes)

	// var x []int = []int{1, 2, 3}
	// x = append(x, 4)
	// x = append(x, 5)
	// x = append(x, 6)
	// x = append(x, x...)
	// x = append(x, []int{7, 8, 9}...)
	// fmt.Println(x)

	var x []int = []int{1, 2, 3}
	// x = appentInt(x, 4, 5)
	// x = appentInt(x, 6, 7)
	copy(x[1:], x[2:])
	fmt.Println(x)
}

func appentInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*cap(x) {
			zcap = 2 * cap(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
