package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://golang.org/",
	"https://bitbucket.org/layer3tv/",
	"https://www.rust-lang.org/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func HandleAsyncCalls() {
	results := asyncHTTPGets(urls)
	for _, result := range results {
		fmt.Printf("%s status %s\n", result.url, result.response.Status)
	}
}


func asyncHTTPGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered channel
	var responses []*HttpResponse
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}
