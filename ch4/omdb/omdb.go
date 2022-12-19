package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Omdb struct {
	Title  string
	Poster string
}

func main() {
	for {
		fmt.Println("请输入电影名称:")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		url := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=cc7e8a19", input)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var omdb Omdb
		json.Unmarshal(body, &omdb)
		download(omdb.Poster)

		resp.Body.Close()
	}
}

func download(imgUrl string) {
	imgPath := "./"
	fileName := path.Base(imgUrl)

	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer resp.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(resp.Body, 32*1024)

	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	// 获取文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d\n", written)
}
