package pubsub

import (
	"fmt"
	"testing"
	"time"
)

func TestMq_Publish(t *testing.T) {
	// New
	mque := New()

	// Subscribe topic
	mque.Subscribe("Notify", func(m *Message) {
		fmt.Println("Received, Topic: Notify")
	})

	// Subscribe topic
	mque.Subscribe("Test2", func(m *Message) {
		fmt.Println("Received, Topic: Test2")
	})

	// Publish value into topic
	mque.Publish(
		Message{
			Topic: "Notify",
			Value: "blablabla",
		},
	)

	// ;)
	time.Sleep(2 * time.Second)

	// Publish value into topic
	mque.Publish(
		Message{
			Topic: "Test2",
			Value: "blablabla",
		},
	)

	// Publish value into topic
	mque.Publish(
		Message{
			Topic: "NotFound",
			Value: "blablabla",
		},
	)

	// ;)
	time.Sleep(1 * time.Second)
}
