package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	GetEndpoints()
}

func GetEndpoints() {
	doc, err := goquery.NewDocument("https://www.conoha.jp/docs/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find("td a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		name := s.Text()

		fmt.Println(url)
		fmt.Println(name)
	})
}
