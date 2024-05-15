package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type server struct {
	wg         sync.WaitGroup
	listener   net.Listener
	shutdown   chan struct{}
	connection chan net.Conn
}

func newServer(address string) (*server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on address %s: %w", address, err)
	}

	return &server{
		listener:   listener,
		shutdown:   make(chan struct{}),
		connection: make(chan net.Conn),
	}, nil
}

func (s *server) acceptConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				continue
			}
			s.connection <- conn
		}
	}
}

func (s *server) handleConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		case conn := <-s.connection:
			go s.handleConnection(conn)
		}
	}
}

func (s *server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Add your logic for handling incoming connections here
	//fmt.Fprintf(conn, "Welcome to my TCP server!\n")
	//time.Sleep(5 * time.Second)
	//fmt.Fprintf(conn, "Goodbye!\n")
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// Print the message to the console.
	fmt.Println("Received data:", string(buf[:reqLen]))

	// Write a response back to the client.
	conn.Write([]byte("Message received."))
}

func (s *server) Start() {
	s.wg.Add(2)
	go s.acceptConnections()
	go s.handleConnections()
}

func (s *server) Stop() {
	close(s.shutdown)
	s.listener.Close()

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-time.After(time.Second):
		fmt.Println("Timed out waiting for connections to finish.")
		return
	}
}

func main() {
	s, err := newServer(":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s.Start()

	// Wait for a SIGINT or SIGTERM signal to gracefully shut down the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Shutting down server...")
	s.Stop()
	fmt.Println("Server stopped.")
	//listen, err := net.Listen("tcp", "localhost:8081")
	//if err != nil {
	//	fmt.Println("Error: ", err.Error())
	//	return
	//}
	//// handle connections in the next steps
	//defer listen.Close()
	//
	//for {
	//	connection, err := listen.Accept()
	//	if err != nil {
	//		fmt.Println("Error: ", err.Error())
	//		return
	//	}
	//	go handleRequest(connection)
	//}
}

func handleRequest(conn net.Conn) {
	//we make a buffer to hold the incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// Print the message to the console.
	fmt.Println("Received data:", string(buf[:reqLen]))

	// Write a response back to the client.
	conn.Write([]byte("Message received."))

	// Close the connection when you're done with it.
	conn.Close()
}
