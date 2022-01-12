package main

import (
	"fmt"
	"log"
	"os"

	"github.com/demidovich/go-learn/02_url_extractor/parser"
	"github.com/demidovich/go-learn/util"
)

func main() {
	targetUrl := targetUrl()

	fmt.Println("Request " + targetUrl)
	html := html(targetUrl)
	urls := parser.Urls(html, targetUrl)

	printInfo(urls)
}

func printInfo(urls []string) {
	fmt.Printf("Total URLs found: %d\n\n", len(urls))
	for _, value := range urls {
		fmt.Println(value)
	}
}

func html(targetUrl string) string {
	html, err := util.HttpGetRequest(targetUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	return html
}

func targetUrl() string {
	if len(os.Args) < 2 {
		log.Fatal("Dont exist URL argument")
	}

	value := os.Args[1]

	if util.IsNotUrl(value) {
		log.Fatal("Invalid URL argument: " + value)
	}

	return value
}
