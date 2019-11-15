package concurrency

import (
	"context"
	"log"
)

func Greeter(ctx context.Context, names <-chan []byte, greetings chan<- []byte, errs <-chan error) {
	for {
		select {
		case <-ctx.Done():
			log.Panicln("finished")
			return
		case err := <-errs:
			log.Println("an error occurred:", err)
		case name := <-names:
			//log.Println("received name: ", name)
			greeting := "Hello " + string(name)
			greetings <- []byte(greeting)
		}
	}
}
