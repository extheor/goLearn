package main

import (
	"bufio"
	"fmt"
	"golearn/ch2/unitconv"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	// 如果缺省命令行参数
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			conv(input.Text())
		}
	} else {
		for _, arg := range args {
			conv(arg)
		}
	}

}

func conv(arg string) {
	// 如果参数为空，则不执行以下逻辑
	isExist := len(arg) > 0
	if !isExist {
		return
	}

	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	f := unitconv.Foot(t)
	m := unitconv.Meter(t)
	p := unitconv.Pound(t)
	k := unitconv.Kilogram(t)

	fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToM(f), m, unitconv.MToF(m))
	fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PToK(p), k, unitconv.KToP(k))
}
