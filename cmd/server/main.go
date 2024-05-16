package main

import (
	"fmt"
	"log"
	"net"
	"sync"
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
	shortServer()

	//s, err := newServer(":8081")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//s.Start()
	//
	//// Wait for a SIGINT or SIGTERM signal to gracefully shut down the server
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//<-sigChan
	//
	//fmt.Println("Shutting down server...")
	//s.Stop()
	//fmt.Println("Server stopped.")

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

func shortServer() {
	newConns := make(chan net.Conn, 128)
	deadConns := make(chan net.Conn, 128)
	publishes := make(chan []byte, 128)
	conns := make(map[net.Conn]bool)

	// a port number will be automatically chosen if add is something like:
	// addr := "127.0.0.1:"
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			newConns <- conn
		}
	}()
	for {
		select {
		case conn := <-newConns:
			conns[conn] = true
			go func() {
				buf := make([]byte, 1024)
				for {
					//conn.SetReadDeadline(time.Now().Add(5 *time.Second))
					nbyte, err := conn.Read(buf)
					if err != nil {
						deadConns <- conn
						break
					} else {
						fragment := make([]byte, nbyte)
						copy(fragment, buf[:nbyte])
						publishes <- fragment
						log.Println("dispatching fragment")
					}
				}
			}()
		case deadConn := <-deadConns:
			_ = deadConn.Close()
			log.Println("connection closed...")
			delete(conns, deadConn)
		case publish := <-publishes:
			for conn, _ := range conns {
				go func(conn net.Conn) {
					totalWritten := 0
					for totalWritten < len(publish) {
						writtenThisCall, err := conn.Write(publish[totalWritten:])
						if err != nil {
							deadConns <- conn
							break
						}
						totalWritten += writtenThisCall
					}
				}(conn)
			}
		}
	}
	listener.Close()
}
