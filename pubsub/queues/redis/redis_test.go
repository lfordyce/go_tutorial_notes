package redis

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/lfordyce/generalNotes/pubsub"
	"log"
	"sync"
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

func _TestTransportReceive(t *testing.T) {

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

func TestSubscriber(t *testing.T) {
	run, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer run.Close()

	client := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         run.Addr(),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	msgToReceive := []byte("hello vice")
	transport := New(WithClient(client))

	var wg sync.WaitGroup
	doneChan := make(chan struct{})

	waitChan := make(chan struct{})
	var once sync.Once
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer close(doneChan)
		for {
			select {
			case <-transport.Done():
				return
			case err := <-transport.ErrChan():
				fmt.Println(err)
				//is.NoErr(err)
				wg.Done()
			case msg := <-transport.Receive("test_receive"):
				//is.Equal(msg, msgToReceive)
				fmt.Println(string(msg))
				wg.Done()
			case <-time.After(2 * time.Second):
				//is.Fail() // time out: transport.Receive
				fmt.Println("timeout reached")
				wg.Done()
				t.Fatal("timeout error")
			default:
				once.Do(func() {
					close(waitChan)
				})
			}
		}
	}(&wg)

	<-waitChan
	cmd, err := run.RPush("test_receive", string(msgToReceive))
	fmt.Println(cmd)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	transport.Stop()
	<-doneChan
}

func TestTransport(t *testing.T) {

	run, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer run.Close()
	//
	//client := redis.NewClient(&redis.Options{
	//	Network:    "tcp",
	//	Addr:       run.Addr(),
	//	Password:   "",
	//	DB:         0,
	//	MaxRetries: 0,
	//})
	client := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         run.Addr(),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	newT := func() pubsub.Transport {
		return New(WithClient(client))
	}
	TransportT(t, newT)
}

func TransportT(t *testing.T, transport func() pubsub.Transport) {
	t.Run("testStandardTransportBehaviour", func(t *testing.T) {
		testStandardTransportBehaviour(t, transport)
	})
}

func testStandardTransportBehaviour(t *testing.T, newTransport func() pubsub.Transport) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("panic recovery fail") // old messages may have confused test
		}
	}()

	transport := newTransport()
	transport1 := newTransport()
	transport2 := newTransport()

	doneChan := make(chan struct{})
	messages := make(map[string][][]byte)
	var wg sync.WaitGroup

	go func() {
		defer close(doneChan)
		for {
			select {
			case <-transport.Done():
				return

			case err := <-transport.ErrChan():
				t.Log(err)
				return

			// test local load balancing with the same transport
			case msg := <-transport.Receive("vicechannel1"):
				messages["vicechannel1"] = append(messages["vicechannel1"], msg)
				wg.Done()
			case msg := <-transport.Receive("vicechannel1"):
				messages["vicechannel1"] = append(messages["vicechannel1"], msg)
				wg.Done()
			case msg := <-transport.Receive("vicechannel1"):
				messages["vicechannel1"] = append(messages["vicechannel1"], msg)
				wg.Done()

			case msg := <-transport.Receive("vicechannel2"):
				messages["vicechannel2"] = append(messages["vicechannel2"], msg)
				wg.Done()
			case msg := <-transport.Receive("vicechannel2"):
				messages["vicechannel2"] = append(messages["vicechannel2"], msg)
				wg.Done()

			case msg := <-transport.Receive("vicechannel3"):
				messages["vicechannel3"] = append(messages["vicechannel3"], msg)
				wg.Done()

			// test distibuted load balancing
			case msg := <-transport.Receive("vicechannel4"):
				messages["vicechannel4.1"] = append(messages["vicechannel4.1"], msg)
				wg.Done()
			case msg := <-transport1.Receive("vicechannel4"):
				messages["vicechannel4.2"] = append(messages["vicechannel4.2"], msg)
				wg.Done()
			case msg := <-transport2.Receive("vicechannel4"):
				messages["vicechannel4.3"] = append(messages["vicechannel4.3"], msg)
				wg.Done()
			}
		}
	}()

	// Let's give some time to initialize all receiving channels
	time.Sleep(time.Millisecond * 10)

	// send 100 messages down each chan
	for i := 0; i < 100; i++ {
		wg.Add(4)
		msg := []byte(fmt.Sprintf("message %d", i+1))
		transport.Send("vicechannel1") <- msg
		transport.Send("vicechannel2") <- msg
		transport.Send("vicechannel3") <- msg
		transport.Send("vicechannel4") <- msg
	}

	wg.Wait()
	transport.Stop()
	transport1.Stop()
	transport2.Stop()
	<-doneChan

	if len(messages) != 6 {
		t.Errorf("expected %d number of messages: actual: %d", 6, len(messages))
	}

	if len(messages["vicechannel1"]) != 100 {
		t.Errorf("expected %d number of messages: actual: %d", 100, len(messages["vicechannel1"]))
	}
	if len(messages["vicechannel2"]) != 100 {
		t.Errorf("expected %d number of messages: actual: %d", 100, len(messages["vicechannel2"]))
	}
	if len(messages["vicechannel2"]) != 100 {
		t.Errorf("expected %d number of messages: actual: %d", 100, len(messages["vicechannel2"]))
	}
	if len(messages["vicechannel4.1"]) == 100 {
		t.Errorf("expected %d number of messages: actual: %d", 100, len(messages["vicechannel4.1"]))
	}
	if len(messages["vicechannel4.2"]) == 100 {
		t.Errorf("expected %d number of messages: actual: %d", 100, len(messages["vicechannel4.2"]))
	}
	if len(messages["vicechannel4.3"]) == 100 {
		t.Errorf("expected %d number of messages: actual: %d", 100, len(messages["vicechannel4.3"]))
	}

	// 	is.Equal(len(messages["vicechannel4.1"])+len(messages["vicechannel4.2"])+len(messages["vicechannel4.3"]), 100)
}
