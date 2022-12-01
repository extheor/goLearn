package popcount

// pc[i] is the population count of i
// var pc [256]byte

// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// PopCount returns the population count (number of set bits) of x.
// func PopCount(x uint64) int {
// 	return int(pc[byte(x>>(0*8))] +
// 		pc[byte(x>>(1*8))] +
// 		pc[byte(x>>(2*8))] +
// 		pc[byte(x>>(3*8))] +
// 		pc[byte(x>>(4*8))] +
// 		pc[byte(x>>(5*8))] +
// 		pc[byte(x>>(6*8))] +
// 		pc[byte(x>>(7*8))])
// }

//  重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。
// func PopCount(x uint64) int {
// 	var res byte
// 	for i := 0; i < 8; i++ {
// 		res += pc[byte(x>>(i*8))]
// 	}
// 	return int(res)
// }

// 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。
// func PopCount(x uint64) int {
// 	var res byte
// 	for i := 0; i < 64; i++ {
// 		res += byte(x >> i & 1)
// 	}
// 	return int(res)
// }

// 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。
func PopCount(x uint64) int {
	var res int
	for x > 0 {
		x = x & (x - 1)
		res++
	}
	return res
}
