package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Comic struct {
	Num        int    `json: num`
	Title      string `json: title`
	Transcript string `json: transcript`
	Alt        string `json: alt`
	Img        string `json: img`
}

func main() {
	// Make a map to store the comics
	comics := make(map[int]Comic)

	// Download and index the comics
	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var comic Comic
		json.Unmarshal(body, &comic)
		comics[comic.Num] = comic
	}

	// Print usage instructions
	fmt.Println("Enter a search term to find matching xkcd comics:")

	// Read search terms from the command line
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.ToLower(input)
		for _, comic := range comics {
			if strings.Contains(strings.ToLower(comic.Transcript), input) || strings.Contains(strings.ToLower(comic.Alt), input) || strings.Contains(strings.ToLower(comic.Title), input) {
				fmt.Println(comic.Img)
			}
		}
	}
}
