package text_tools

import (
	"fmt"
	"regexp"
)

func Concat(left, right string) string {
	return left + right
}

func FetchUrl(str string) []string {

	var (
		schema = `(https?://)?`
		host   = `([a-z\\d_-]+\.){1,10}[a-z]{2,10}`
		port   = `(:\d{1,5})?`
		path   = `(/[a-zA-Z\d_.-]+)*`
		params = `(\?[^\s$]*)?`
	)

	r, err := regexp.Compile(`(?i)` + schema + host + port + path + params)
	if err != nil {
		fmt.Println(err)
	}

	return r.FindAllString(str, -1)
}
