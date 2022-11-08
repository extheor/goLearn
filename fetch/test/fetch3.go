package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 如果 http:// 前缀不存在
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		fmt.Println("HTTP协议状态码: ", resp.Status)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reader: %s: %v\n", url, err)
		}
	}
}
