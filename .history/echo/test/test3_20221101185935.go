package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now() // 获取当前时间
	startTimestamp := startTime.Unix() // 时间戳
	fmt.Println(startTimestamp)

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	endTime := time.Now() // 获取当前时间
	endTimestamp := endTime.Unix() // 时间戳
	fmt.Println(endTimestamp)

	fmt.Println(endTimestamp - startTimestamp)
}