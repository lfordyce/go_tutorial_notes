package concurrency

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 50               // number of Poller goroutines to launch
	pollInterval   = 10 * time.Second // how often to poll each URL
	statusInterval = 10 * time.Second // how often to log status to stdout
	errTimeout     = 10 * time.Second // back-off timeout on error
)

var urls2 = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// State represents the last-known state of a URL.
type State struct {
	url    string
	status string
	date   string
}

// StateMonitor maintains a map that stores the state of the URLs being
// polled, and prints the current state every updateInterval nanoseconds.
// It returns a chan State to which resource state should be sent.
func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]State)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s
			}
		}
	}()
	return updates
}

// logState prints a state map.
func logState(s map[string]State) {
	log.Println("Current state:")
	for k, v := range s {
		log.Printf(" url=[%s] date=[%s] status=[%s]", k, v.date, v.status)
	}
}

// Resource represents an HTTP URL to be polled by this program.
type Resource struct {
	url      string
	errCount int
}

// Poll executes an HTTP HEAD request for url
// and returns the HTTP status string or an error string.
func (r *Resource) Poll() (header *http.Response, err error) {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return nil, err
	}
	r.errCount = 0
	return resp, nil
}

// Sleep sleeps for an appropriate interval (dependent on error state)
// before sending the Resource to done.
func (r *Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s, _ := r.Poll()
		date := s.Header.Get("Date")
		status <- State{r.url, s.Status, date}
		out <- r
	}
}
