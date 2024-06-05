package main

import (
	"flag"
	"fmt"
	"github.com/lfordyce/generalNotes/cmd"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	port = flag.Int("port", 8080, "port to listen on")
)

type Server struct {
	conn     *net.UDPConn
	messages chan string
	clients  map[*uuid.UUID]Client
}

type Client struct {
	userID   uuid.UUID
	userName string
	userAddr *net.UDPAddr
}

type Message struct {
	messageType      cmd.MessageType
	userID           *uuid.UUID
	userName         string
	content          string
	connectionStatus cmd.ConnectionStatus
	time             string
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func (s *Server) handleMessage() {
	var buf [512]byte

	n, addr, err := s.conn.ReadFromUDP(buf[:])
	if err != nil {
		return
	}
	fragment := make([]byte, n)
	copy(fragment, buf[:n])

	msg := string(fragment)
	m := s.parseMessage(msg)

	if m.connectionStatus == cmd.LEAVING {
		delete(s.clients, m.userID)
		s.messages <- msg
		log.Printf("%s left", m.userName)
	} else {
		switch m.messageType {
		case cmd.FUNC:
			c := Client{
				userID:   *m.userID,
				userName: m.userName,
				userAddr: addr,
			}
			s.clients[m.userID] = c
			s.messages <- msg
			log.Printf("%s joining", m.userName)
		case cmd.CLASSIQUE:

			log.Printf("%s %s: %s", m.time, m.userName, m.content)
			s.messages <- msg
		}
	}
}

func (s *Server) parseMessage(msg string) (m Message) {
	stringArray := strings.Split(msg, "\x01")

	parsedUUID, _ := uuid.Parse(stringArray[0])
	m.userID = &parsedUUID
	messageTypeStr, _ := strconv.Atoi(stringArray[1])
	m.messageType = cmd.MessageType(messageTypeStr)
	m.userName = stringArray[2]
	m.content = stringArray[3]
	m.time = stringArray[4]
	// pf("MESSAGE RECEIVED: %s \n", msg)
	// pf("USER NAME: %s \n", stringArray [2])
	// pf("CONTENT: %s \n", stringArray [3])
	if strings.HasPrefix(msg, ":q") || strings.HasPrefix(msg, ":quit") {
		log.Printf("%s is leaving\n", m.userName)
		m.connectionStatus = cmd.LEAVING
	}
	return
}

func (s *Server) sendMessage() {
	for {
		msg := <-s.messages
		//p(00, sendstr)
		for _, c := range s.clients {
			n, err := s.conn.WriteToUDP([]byte(msg), c.userAddr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
				os.Exit(1)
			}
			fmt.Printf("Bytes read %d, error: %v\n", n, err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	//arguments := os.Args
	//if len(arguments) == 1 {
	//	fmt.Println("Please provide port number")
	//	return
	//}
	//
	//PORT := ":" + arguments[1]

	//pc, err := net.ListenPacket("udp", fmt.Sprintf(":%d", *port))
	//if err != nil {
	//	panic(err)
	//}
	//defer pc.Close()
	//
	//for {
	//	buf := make([]byte, 1024)
	//	n, addr, err := pc.ReadFrom(buf)
	//	if err != nil {
	//		continue
	//	}
	//	go serve(pc, addr, buf[:n])
	//}

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}

	net.ListenPacket("udp", fmt.Sprintf(":%d", *port))

	svc := &Server{
		messages: make(chan string, 10),
		clients:  make(map[*uuid.UUID]Client),
	}
	svc.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
	defer svc.conn.Close()
	go svc.sendMessage()

	for {
		svc.handleMessage()
	}

	//buffer := make([]byte, 1024)
	//rand.Seed(time.Now().Unix())
	//for {
	//	n, addr, err := conn.ReadFromUDP(buffer)
	//	fmt.Print("-> ", string(buffer[0:n-1]))
	//
	//	if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
	//		fmt.Println("Exiting UDP server!")
	//		return
	//	}
	//
	//	data := []byte(strconv.Itoa(random(1, 1001)))
	//	fmt.Printf("data: %s\n", string(data))
	//	_, err = conn.WriteToUDP(data, addr)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	pc.WriteTo(buf, addr)
}
