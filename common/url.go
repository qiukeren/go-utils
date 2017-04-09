package common

import (
	"errors"
	"net/url"
	"strings"
)

//urlencode
func UrlEncode(url1 string) string {
	return url.QueryEscape(url1)
}

func UrlDecode(url1 string) (string, error) {
	return url.QueryUnescape(url1)
}

//FormatUrl is a function for spiders to format urls when the url comes like "/path/to/html" to "http://www.xxx.com/path/to/html"
func FormatUrl(url1, site string) (string, error) {

	if strings.HasPrefix(strings.TrimSpace(url1), "javascript:") {
		return "", errors.New("starts with 'javascript:'")
	}

	u, err := url.Parse(url1)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "http"

	}
	if u.Host == "" {
		u.Host = site
	}
	return u.String(), nil
}

//IsCurrentSite is a function to test the current url belongs to the current site.
func IsCurrentSite(url1 string, site string, protocol string) bool {
	u, err := url.Parse(url1)
	if err != nil {
		return false
	}
	if u.Scheme != protocol {
		return false
	}
	if u.Host == site || u.Host == "" {
		return true
	}
	return false
}
