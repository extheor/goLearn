package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now() // 获取当前时间
	startTimestamp := startTime.UnixNano() // 时间戳

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(strings.Join(os.Args, " "))

	endTime := time.Now() // 获取当前时间
	endTimestamp := endTime.UnixNano() // 时间戳

	fmt.Println(endTimestamp - startTimestamp)
}