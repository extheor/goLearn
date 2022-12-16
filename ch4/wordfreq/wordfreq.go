package main

import (
	"bufio"
	"fmt"
	"os"
)

// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
func main() {
	mapv := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		str := input.Text()
		mapv[str]++

		fmt.Printf("The length %d\n", len(mapv))
		for k, v := range mapv {
			fmt.Printf("%s %d\n", k, v)
		}
	}
}
