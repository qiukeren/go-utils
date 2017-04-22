package common

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"strings"
)

//Get is a common http client for httpGet operations and hide UA as GoogleBot.
func Get(urls string) ([]byte, error) {

	req, err := http.NewRequest("GET", urls, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Googlebot/2.1 (+http://www.google.com/bot.html)")

	client := getClient()
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	tempData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err

	}

	return tempData, nil
}

func getClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		DisableKeepAlives:  true,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
	return client
}

func RequestWithHeader(
	method string,
	url1 string,
	timeout int,
	header map[string]string,
	params map[string]string,
) ([]byte, error) {

	client := getClient()

	paramsA := ""
	if method == "POST" {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		paramsA = values.Encode()
	}

	req, _ := http.NewRequest(method, url1, strings.NewReader(paramsA))

	for k, v := range header {
		req.Header.Set(k, v)
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if method == "GET" {
		values := req.URL.Query()
		for k, v := range params {
			values.Add(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	getResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer getResp.Body.Close()
	body, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
