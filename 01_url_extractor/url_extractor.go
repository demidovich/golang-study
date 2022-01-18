package main

import (
	"fmt"
	"log"
	netUrl "net/url"
	"os"

	"github.com/demidovich/golang-study/01_url_extractor/parser"
	"github.com/demidovich/golang-study/pkg/util"
)

func main() {
	targetUrl := targetUrl()

	fmt.Println("Request " + targetUrl)
	html := html(targetUrl)
	host := hostFromUrl(targetUrl)
	urls := parser.Urls(html, host)

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

func hostFromUrl(url string) string {
	u, err := netUrl.Parse(url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		log.Fatal("Error extract Host from targetUrl: " + url)
	}
	return u.Scheme + "://" + u.Host
}
