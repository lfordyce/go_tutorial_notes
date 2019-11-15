package utils

import (
	"fmt"
	"time"
)

var data = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

type Harvester struct {
	ticker *time.Ticker
	add    chan string
	urls   []string
}

func NewHarvester() *Harvester {
	rv := &Harvester{
		ticker: time.NewTicker(time.Second * 3),
		add:    make(chan string),
		urls:   data,
	}
	go rv.run()
	return rv
}

func (h *Harvester) Execute() {

}

func (h *Harvester) run() {
	for {
		select {
		case <-h.ticker.C:
			for _, u := range h.urls {
				fmt.Println(u)
			}
		case u := <-h.add:
			h.urls = append(h.urls, u)
		}
	}
}

func (h *Harvester) AddURL(u string) {
	h.add <- u
}
