package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	port = flag.Int("port", 8080, "port to listen on")
)

type Server struct {
	conn     *net.UDPConn
	messages chan string
	clients  map[int64]Client
}

type Client struct {
	userID   int64
	userName string
	userAddr *net.UDPAddr
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
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

	s, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	pc.WriteTo(buf, addr)
}
