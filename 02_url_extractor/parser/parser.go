package parser

import (
	"regexp"
	"strings"
)

func Urls(html string, host string) []string {
	host = strings.TrimSuffix(host, "/")

	reg, _ := regexp.Compile(`(?i)\s+href\=(?:['"]{1})?([^'"\s<]+)`)
	matches := reg.FindAllStringSubmatch(html, -1)
	if matches == nil {
		return make([]string, 0)
	}

	results := make([]string, 0, len(matches))

	for _, match := range matches {
		if isLink(match[1]) {
			results = append(
				results,
				urlWithHost(match[1], host),
			)
		}
	}

	return results
}

func urlWithHost(url string, host string) string {
	reg, _ := regexp.Compile(`(?i)^https?\://`)
	if reg.MatchString(url) {
		return url
	}

	if strings.HasPrefix(url, "/") {
		return host + url
	}

	return host + "/" + url
}

func isLink(url string) bool {
	regJs, _ := regexp.Compile(`(?i)^javascript:`)
	if regJs.MatchString(url) {
		return false
	}

	regLink, _ := regexp.Compile(`^(.+/)?[^.]+(\.(php|htm?l|asp))?(\?.*)?$`)
	return regLink.MatchString(url)
}
