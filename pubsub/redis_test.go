package pubsub

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"testing"
	"time"
)

// Greeter is a service that greets people.
func Greeter(ctx context.Context, names <-chan []byte, greetings chan<- []byte, errs <-chan error) {
	for {
		select {
		case <-ctx.Done():
			log.Println("finished")
			return
		case err := <-errs:
			log.Println("an error occurred:", err)
		case name := <-names:
			greeting := "Hello " + string(name)
			greetings <- []byte(greeting)
		}
	}
}

func TestTransportReceive(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	transport := New(WithClient(client))
	defer func() {
		transport.Stop()
		<-transport.Done()
	}()

	names := transport.Receive("names")
	greetings := transport.Send("greetings")
	Greeter(ctx, names, greetings, transport.ErrChan())

	//greetings <- []byte("batman")
	//go func() {
	//	select {
	//	case name := <-names:
	//		fmt.Println(name)
	//	}
	//}()

	//select {
	//case <-time.After(1 * time.Second):
	//	greetings <- []byte("batman")
	//case name := <-names:
	//	fmt.Println(name)
	//}
}
