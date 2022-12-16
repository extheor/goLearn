package main

import (
	"golearn/ch4/github"
	"os"
)

func main() {
	github.IssueCLI(os.Args[len(os.Args)-1])
}
