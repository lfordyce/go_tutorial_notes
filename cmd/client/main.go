package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/lfordyce/generalNotes/cmd"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	connection          *net.UDPConn
	alive               bool
	userID              uuid.UUID
	userName            string
	sendingMessageQueue chan string
	receiveMessages     chan string
}

var scanError error

func (c *Client) packMessage(msg string, messageType cmd.MessageType) string {
	return strings.Join([]string{c.userID.String(), strconv.Itoa(int(messageType)), c.userName, msg, time.Now().Format("15:04:05")}, "\x01")
}

func (c *Client) funcSendMessage(msg string) {
	message := c.packMessage(msg, cmd.FUNC)
	_, err := c.connection.Write([]byte(message))
	checkError(err, "func_sendMessage")
}

func (c *Client) sendMessage() {
	for c.alive {

		msg := <-c.sendingMessageQueue
		message := c.packMessage(msg, cmd.CLASSIQUE)
		_, err := c.connection.Write([]byte(message))
		checkError(err, "sendMessage")
	}
}

func (c *Client) receiveMessage() {
	var buf [512]byte
	//var userID *uuid.UUID
	for c.alive {
		n, err := c.connection.Read(buf[0:])
		checkError(err, "receiveMessage")
		//msg := string(buf[0:n])
		//stringArray := strings.Split(msg, "\x01")

		//userID, err = uuid.ParseHex(stringArray[0])
		//checkError(err, "receiveMessage")
		//if *userID != c.userID {
		c.receiveMessages <- string(buf[0:n])
		fmt.Println("")
		//}
	}
}

func (c *Client) readInput() {
	var msg string
	for c.alive {
		fmt.Println("msg: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			msg = scanner.Text()
			if msg == ":quit" || msg == ":q" {
				c.alive = false
			}
			c.sendingMessageQueue <- msg
		}
		//_,scanError := fmt.Scanln(&msg)
		//checkError(scanError, "readInput")
	}
}

func (c *Client) printMessage() {
	for c.alive {
		msg := <-c.receiveMessages
		stringArray := strings.Split(msg, "\x01")
		var userName = stringArray[2]
		var content = stringArray[3]
		var timestamp = stringArray[4]
		fmt.Printf("%s %s: %s", timestamp, userName, content)
		fmt.Println("")
		// pf("MESSAGE RECEIVED: %s \n", msg)
		// pf("USER NAME: %s \n", stringArray [2])
		// pf("CONTENT: %s \n", stringArray [3])
		if strings.HasPrefix(msg, ":q") || strings.HasPrefix(msg, ":quit") {
			fmt.Printf("%s is leaving", userName)
		}
	}
}

func nowTime() string {
	return time.Now().String()
}
func checkError(err error, funcName string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s-----in func:%s", err.Error(), funcName)
		os.Exit(1)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", CONNECT)
	checkError(err, "main")
	c := &Client{
		alive:               true,
		sendingMessageQueue: make(chan string),
		receiveMessages:     make(chan string),
	}
	//uuid.New()
	//u, err := uuid.NewV4()
	//checkError(err, "main")
	c.userID = uuid.New()

	fmt.Println("input name: ")
	_, err = fmt.Scanln(&c.userName)
	checkError(err, "main")

	c.connection, err = net.DialUDP("udp", nil, udpAddr)
	checkError(err, "main")
	defer c.connection.Close()

	c.funcSendMessage("joined")

	go c.printMessage()
	go c.receiveMessage()

	go c.sendMessage()
	c.readInput()

	c.funcSendMessage("left")

	os.Exit(0)

	//c, err := net.DialUDP("udp4", nil, s)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	//defer c.Close()
	//
	//for {
	//	reader := bufio.NewReader(os.Stdin)
	//	fmt.Print(">> ")
	//	text, _ := reader.ReadString('\n')
	//	data := []byte(text + "\n")
	//	_, err = c.Write(data)
	//	if strings.TrimSpace(string(data)) == "STOP" {
	//		fmt.Println("Exiting UDP client!")
	//		return
	//	}
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	buffer := make([]byte, 1024)
	//	n, _, err := c.ReadFromUDP(buffer)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	//}
}
