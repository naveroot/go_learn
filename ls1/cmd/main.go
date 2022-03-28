package main

import (
	"flag"
	"fmt"
	"ls1/pkg/crawler"
	"ls1/pkg/crawler/spider"
	"strings"
)

func main() {
	urls := []string{"https://golang.org", "https://go.dev/"}
	gs := gosearch{}
	gs.scanner = spider.New()
	chRes, _ := gs.scanner.BatchScan(urls, 3, 3)
	docs := []crawler.Document{}

	for rec := range chRes {
		docs = append(docs, rec)
	}


	var keyword string
	flag.StringVar(&keyword, "s", "", "Search word")
	flag.Parse()
	fmt.Println("Search result by keyworld:", keyword)
	if keyword != "" {
		for _, doc := range docs {
			if strings.Contains(doc.Title, keyword) {
				fmt.Println(doc.URL, doc.Title)
			}
		}
	}
}

type gosearch struct {
	scanner crawler.Interface
}
