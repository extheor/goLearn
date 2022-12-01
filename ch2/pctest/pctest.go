package main

import (
	"fmt"
	"golearn/ch2/popcount"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(popcount.PopCount(11111))
	secs := time.Since(start).Seconds()
	fmt.Printf("执行时间: %g", secs)
}
