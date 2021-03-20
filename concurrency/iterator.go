package concurrency

import "strings"

type Iterate interface {
	Next() (interface{}, bool)
	Error() error
}

type iterator struct {
	valueChan <-chan interface{}
	okChan    <-chan bool
	errChan   <-chan error
	err       error
}

func (i *iterator) Next() (interface{}, bool) {
	var (
		value interface{}
		ok    bool
	)
	value, ok, i.err = <-i.valueChan, <-i.okChan, <-i.errChan
	return value, ok
}

func (i *iterator) Error() error {
	return i.err
}

// Generator function that produces data
func NewIterator(data []string) Iterate {
	out := make(chan interface{})
	ok := make(chan bool)
	err := make(chan error)
	// Go Routine
	go func() {
		defer close(out) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				out <- word // Send word to channel and waits for its reading
				ok <- true
				err <- nil // if there was any error, change its value
			}
		}
		out <- ""
		ok <- false
		err <- nil
	}()

	return &iterator{out, ok, err, nil}
}
