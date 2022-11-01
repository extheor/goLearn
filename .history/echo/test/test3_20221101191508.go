package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now() // 获取当前时间
	startTimestamp := startTime.Unix() // 时间戳

	// var s, sep string
	// for i := 0; i < 100000; i++ {
	// 	s += sep + string(rune(i))
	// 	sep = " "
	// }
	fmt.Println(strings.Join(os.Args, " "))

	endTime := time.Now() // 获取当前时间
	endTimestamp := endTime.Unix() // 时间戳

	fmt.Println(endTimestamp - startTimestamp)
}