package github

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

var title = flag.String("title", "", "issues title")
var body = flag.String("body", "", "issues body")
var number = flag.Int("number", 1, "issues number")

func IssueCLI(term string) {
	flag.Parse()

	params := make(map[string]string)
	params["title"] = *title
	params["body"] = *body
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	if term == "create" {
		createStatus := map[string]int{
			"Created":            201,
			"Forbidden":          403,
			"Resource not found": 404,
			"Gone":               410,
			"Validation failed, or the endpoint has been spammed.": 422,
			"Service unavailable": 503,
		}

		req, err := http.NewRequest("POST", BaseIssuesURL, bytes.NewBuffer(bytesData))
		if err != nil {
			fmt.Printf("Error creating")
		}

		req.Header.Set("Content-Type", "application/vnd.github+json")
		req.Header.Set("Authorization", "Bearer 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")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		if resp.StatusCode == createStatus["Created"] {
			fmt.Println("创建成功")
		} else {
			resp.Body.Close()

			for i, v := range createStatus {
				if resp.StatusCode == v {
					fmt.Printf("create issue failed: %s %s", i, resp.Status)
				}
			}
		}

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		fmt.Println("创建成功")
		// fmt.Println(string(content))

		resp.Body.Close()
	} else if term == "get" {
		getStatus := map[string]int{
			"OK":                 200,
			"Moved permanently":  301,
			"Not modified":       304,
			"Resource not found": 404,
			"Gone":               410,
		}

		url := BaseIssuesURL + "/" + strconv.Itoa(*number)
		req, err := http.NewRequest("GET", url, bytes.NewBuffer(bytesData))
		if err != nil {
			fmt.Printf("Error geting")
		}

		req.Header.Set("Content-Type", "application/vnd.github+json")
		req.Header.Set("Authorization", "Bearer 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")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		if resp.StatusCode == getStatus["OK"] {
			fmt.Println("读取成功")
		} else {
			resp.Body.Close()

			for i, v := range getStatus {
				if resp.StatusCode == v {
					fmt.Printf("get issue failed: %s %s", i, resp.Status)
				}
			}
		}

		var result Issue
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			fmt.Printf("decode failed: %s", err)
		}

		fmt.Println(result)

		resp.Body.Close()
	} else if term == "update" {
		updateStatus := map[string]int{
			"OK":                 200,
			"Moved permanently":  301,
			"Forbidden":          403,
			"Resource not found": 404,
			"Gone":               410,
			"Validation failed, or the endpoint has been spammed.": 422,
			"Service unavailable": 503,
		}

		url := BaseIssuesURL + "/" + strconv.Itoa(*number)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bytesData))
		if err != nil {
			fmt.Printf("Error geting")
		}

		req.Header.Set("Content-Type", "application/vnd.github+json")
		req.Header.Set("Authorization", "Bearer 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")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		if resp.StatusCode == updateStatus["OK"] {
			fmt.Println("更新成功")
		} else {
			resp.Body.Close()

			for i, v := range updateStatus {
				if resp.StatusCode == v {
					fmt.Printf("get issue failed: %s %s", i, resp.Status)
				}
			}
		}

		var result Issue
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			fmt.Printf("decode failed: %s", err)
		}

		fmt.Println(result)

		resp.Body.Close()
	} else if term == "unlock" {
		updateStatus := map[string]int{
			"No Content":         204,
			"Forbidden":          403,
			"Resource not found": 404,
		}

		url := BaseIssuesURL + "/" + strconv.Itoa(*number) + "/lock"
		req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(bytesData))
		if err != nil {
			fmt.Printf("Error geting")
		}

		req.Header.Set("Content-Type", "application/vnd.github+json")
		req.Header.Set("Authorization", "Bearer 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")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		if resp.StatusCode == updateStatus["No Content"] {
			fmt.Println("关闭成功")
		} else {
			resp.Body.Close()

			for i, v := range updateStatus {
				if resp.StatusCode == v {
					fmt.Printf("get issue failed: %s %s\n", i, resp.Status)
				}
			}
		}

		resp.Body.Close()
	}
}
