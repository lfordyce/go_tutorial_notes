package concurrency

import (
	"fmt"
	"time"
)

type clientID string
type command string
type topic string

type topicChannel chan clientPair
type clientChannel chan string

type subscriptions map[clientID]clientChannel

type message struct {
	command command
	channel clientChannel
	topic   topic
	client  clientID
}

type clientPair struct {
	ID      clientID
	Channel clientChannel
}

func roundRobin(left chan message, right chan message) chan message {
	messageChan := make(chan message)
	current := false

	go func() {
		for msg := range messageChan {
			switch string(msg.command) {
			case "unsubscribe":
				left <- msg
				right <- msg
			case "subscribe":
				left <- msg
				right <- msg
			case "send":
				switch current {
				case false:
					left <- msg
				default:
					right <- msg
				}
				current = !current
			}
		}
	}()

	return messageChan
}

func dispatcher(topics map[topic]topicChannel) chan message {

	routes := make(map[topic]subscriptions)
	for topic := range topics {
		routes[topic] = make(subscriptions)
	}

	messageChan := make(chan message)

	go func() {
		for msg := range messageChan {
			switch string(msg.command) {
			case "unsubscribe":

				// delete the client from the topic
				delete(routes[msg.topic], msg.client)
			case "subscribe":

				// add the client to the topic
				routes[msg.topic][msg.client] = msg.channel
			case "send":
				for clientID, client := range routes[msg.topic] {
					topics[msg.topic] <- clientPair{ID: clientID, Channel: client}
				}
			}
		}
	}()

	return messageChan
}

func feed(name string, interval time.Duration) topicChannel {

	topicChan := make(topicChannel)

	alreadySeen := make(map[clientID]int)
	var latestData []string

	go func(name string) {

		t := time.NewTicker(interval)

		var messageID int
		for {
			select {
			case <-t.C:
				messageID++
				latestData = append(latestData, fmt.Sprintf("[%s.%d] [%v] message", name, messageID, time.Now().Unix()))
			case clientPair := <-topicChan:

				i, ok := alreadySeen[clientPair.ID]
				if !ok {
					i = 0
				}

				last := i

				for _, data := range latestData[i:] {
					clientPair.Channel <- data
					last++
				}

				alreadySeen[clientPair.ID] = last
			}
		}
	}(name)

	return topicChan
}

func client(name string) clientChannel {
	clientChan := make(clientChannel)

	go func() {
		for {
			select {
			case data := <-clientChan:
				fmt.Printf("[%s] recv %s\n", name, data)
			}
		}
	}()

	return clientChan
}

func execute() {
	client1 := client("client1")
	client2 := client("client2")

	topic1 := feed("feed1", 500*time.Millisecond)
	topic2 := feed("feed2", 700*time.Millisecond)

	topics := make(map[topic]topicChannel)
	topics["feed1"] = topic1
	topics["feed2"] = topic2

	C := roundRobin(dispatcher(topics), dispatcher(topics))

	// create a subscription from client 1 to feed 1
	C <- message{
		command: "subscribe",
		channel: client1,
		topic:   topic("feed1"),
		client:  "client1",
	}

	// create a subscription from client 2 to feed 1
	C <- message{
		command: "subscribe",
		channel: client2,
		topic:   topic("feed1"),
		client:  "client2",
	}

	// create a subscription from client 2 to feed 2
	C <- message{
		command: "subscribe",
		channel: client2,
		topic:   topic("feed2"),
		client:  "client2",
	}

	// wake up the dispatcher
	go func() {
		for {
			<-time.After(100 * time.Millisecond)

			C <- message{
				command: "send",
				topic:   topic("feed1"),
			}
			C <- message{
				command: "send",
				topic:   topic("feed2"),
			}
		}
	}()

	<-time.After(5 * time.Second)

	// unsubscribe client 1 from feed 1 and feed 2
	C <- message{
		command: "unsubscribe",
		channel: client1,
		topic:   topic("feed1"),
		client:  "client1",
	}

	// unsubscribe client 1 from feed 1 and feed 2
	C <- message{
		command: "unsubscribe",
		channel: client1,
		topic:   topic("feed2"),
		client:  "client1",
	}

	<-time.After(5 * time.Second)
}
