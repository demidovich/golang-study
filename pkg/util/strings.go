package util

import (
	"net/url"
)

func IsUrl(value string) bool {
	u, err := url.Parse(value)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsNotUrl(value string) bool {
	return !IsUrl(value)
}
