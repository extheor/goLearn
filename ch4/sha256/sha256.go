package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hashMethod = flag.Int("s", 256, "选择哈希版本：256、384、512。")

func main() {
	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// fmt.Println(c1)
	// fmt.Println(c2)

	// r := findDiff(c1, c2)
	// fmt.Println(r)

	flag.Parse()
	printHash()
}

// 编写一个函数，计算两个SHA256哈希码中不同bit的数目。
func findDiff(c1, c2 [32]byte) int {
	var count int
	for i := 0; i < len(c1); i++ {
		temp := c1[i] ^ c2[i]
		count += popCount(temp)
	}

	return count
}

// 计算一个数字中包含1bit的个数
func popCount(x byte) int {
	var res int
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return res
}

// 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
func printHash() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("输入要解析的字符串：")
	input.Scan()

	switch *hashMethod {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(input.Text())))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(input.Text())))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(input.Text())))
	}
}
