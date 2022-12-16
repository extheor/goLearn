package main

import (
	"fmt"
	"golearn/ch4/github"
	"log"
	"os"
	"time"
)

type timeclass struct {
	不到一个月的 []github.Issue
	不到一年的  []github.Issue
	超过一年的  []github.Issue
}

func main() {
	var timemap timeclass
	curtime := time.Now()
	curyear := curtime.Year()
	curmonth := int(curtime.Month())

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		createat := item.CreatedAt

		// 如果当年大于创建issues时的年份
		createyear := createat.Year()
		createmonth := int(createat.Month())

		diffyear := curyear - createyear

		if diffyearmonth := curmonth - createmonth; diffyear >= 1 || diffyear == 1 && diffyearmonth >= 0 {
			// 超过一年
			timemap.超过一年的 = append(timemap.超过一年的, *item)
		} else if diffyear == 1 && 12+diffyearmonth >= 1 {
			// 超过一个月但是不到一年
			timemap.不到一年的 = append(timemap.不到一年的, *item)
		} else {
			// 不到一个月
			timemap.不到一个月的 = append(timemap.不到一个月的, *item)
		}

		// fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Printf("%#v", timemap)
}
