package main

import (
	"flag"
	"fmt"
	"log"
	"ls1/pkg/crawler"
	"ls1/pkg/crawler/spider"
	"strings"

)

func main() {
	urls := []string{"https://golang.org", "https://go.dev/"}
	var keyword string
	flag.StringVar(&keyword, "s", "", "Search word")
	flag.Parse()
	fmt.Println("Search result by keyworld:", keyword)

	docs, err := gosearch(urls)
	if err != nil {
		log.Fatal(err)
		return
	}

	if keyword != "" {
		for _, doc := range docs {
			if strings.Contains(doc.Title, keyword) {
				fmt.Println(doc.URL, doc.Title)
			}
		}
	}

}


func gosearch(urls []string) ([]crawler.Document, error) {
	var result []crawler.Document
	spdr := spider.New()
	for _,url := range urls {
		docs, err := spdr.Scan(url, 3)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		for _, doc := range docs {
			result = append(result, doc)
		}
	} 
	return result, nil
}