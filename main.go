package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestURL struct {
	url    string
	status string
}

var errRequestFailed = errors.New("requestURL failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestURL)
	//= make(map[string]string)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazone.com",
		"https://www.reddit.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hitURL(url string, c chan<- requestURL) {

	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestURL{url: url, status: status}

}
