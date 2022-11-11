package utils

import (
	"context"
	"fmt"
	"net/url"
	"testing"
)

type server struct {
	url        url.URL
	serverName string
}

func TestLoadBalance(t *testing.T) {
	server1, _ := url.Parse("https://localhost:3000")
	server2, _ := url.Parse("https://localhost:3001")
	server3, _ := url.Parse("https://localhost:3002")

	servers := []server{
		{
			url:        *server1,
			serverName: "server1",
		},
		{
			url:        *server2,
			serverName: "server2",
		},
		{
			url:        *server3,
			serverName: "server3",
		},
	}

	var (
		lb LoadBalance[server]
	)

	// round-robin
	lb, _ = NewRoundRobin(servers...)

	next, err := lb.Next(context.TODO())
	if err != nil {
		return
	}

	fmt.Println("round robin:", next)

	lb, _ = NewLeastConnection(servers...)
	next1, err := lb.Next(context.TODO())
	if err != nil {
		return
	}

	fmt.Println("least connection:", next1)

}
