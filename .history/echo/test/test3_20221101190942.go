package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now() // 获取当前时间
	startTimestamp := startTime.Unix() // 时间戳

	var s, sep string
	for i := 0; i < 1000; i++ {
		s += sep + string(i)
		sep = " "
	}

	endTime := time.Now() // 获取当前时间
	endTimestamp := endTime.Unix() // 时间戳

	fmt.Println(endTimestamp - startTimestamp)
}