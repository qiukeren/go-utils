package common

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
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
