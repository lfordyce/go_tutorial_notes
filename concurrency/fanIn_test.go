package concurrency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestRepeatFunc(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	randFn := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFunc(done, randFn), 10) {
		fmt.Println(num)
	}
}

func TestTeeOperation(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeater(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}

func dataGen(done <-chan struct{}, iter int) <-chan int {
	streamID := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	genID := 0

	go func() {
		defer close(streamID)
		for i := 0; i < iter; i++ {
			genID = r.Intn(100)
			select {
			case <-done:
				return
			case streamID <- genID:
			}
		}
	}()
	return streamID
}

func fanOutFunc(done chan struct{}, in <-chan int) <-chan string {
	resultValue := make(chan string)
	go func() {
		defer close(resultValue)
		for n := range in {
			select {
			case <-done:
				fmt.Println("fanOutFunc has been cancled")
				return
			case resultValue <- parse(done, n) + " _ Processed":
			}
		}
	}()
	return resultValue
}

func fanInn(done <-chan struct{}, cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	multiplex := make(chan string)

	multiplexFunc := func(c <-chan string) {
		defer wg.Done()
		for text := range c {
			select {
			case <-done:
				fmt.Println("funIn has been canceled")
				return
			case multiplex <- text:
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go multiplexFunc(c)
	}

	go func() {
		wg.Wait()
		close(multiplex)
	}()
	return multiplex
}

func TestFanInFanOut(t *testing.T) {
	nWorkers := 4
	nJobs := 8
	done := make(chan struct{})

	fanOut := make([]<-chan string, nWorkers)
	for i := 0; i < nWorkers; i++ {
		fanOut[i] = fanOutFunc(done, dataGen(done, nJobs))
	}

	for result := range fanInn(done, fanOut...) {
		fmt.Println(result)
	}
}

type post struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func parse(done chan struct{}, id int) string {

	body, err := getBody(id)
	//close done when we get an error while parsing
	if err != nil {
		close(done)
		log.Fatal(err)
	}

	var post post

	if err := json.Unmarshal(body, &post); err != nil {
		close(done)
		log.Fatalf("Can't unmarshal: %v", err)
	}

	//longestPost := 0
	//longestPostID := 0
	//longestPostEmail := ""

	//for _, p := range posts {
	//	if len(p.Body) > longestPost {
	//		longestPost = len(p.Body)
	//		longestPostID = p.PostID
	//		longestPostEmail = p.Email
	//	}
	//}

	//return fmt.Sprintf("%d %s %d", longestPostID, longestPostEmail, longestPost)
	return fmt.Sprintf("%d %s %d", post.ID, post.Email, len(post.Body))

}

func getBody(id int) ([]byte, error) {
	site, err := url.Parse("https://jsonplaceholder.typicode.com/comments/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	/*
		q := site.Query()
		q.Set("postId", strconv.Itoa(id))
		site.RawQuery = q.Encode()
		log.Println("Getting: ", site.String())
	*/

	client := &http.Client{}
	req, err := http.NewRequest("GET", site.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
