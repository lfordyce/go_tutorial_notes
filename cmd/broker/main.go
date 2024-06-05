package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ChatMessage struct {
	msg string
	id  int
}

type User struct {
	id       int
	name     string
	muteList map[int]string //users to mute by id
}

// Global vars
var state = []User{}
var userNameList = map[string]int{} // map userNames to user ID
var msgHistory []string

func main() {
	port := flag.Int("port", 3478, "Listening port.")
	flag.Parse()

	// add system user for system notifications
	addNewUser("system")
	// create and start our primary chatroom
	b := NewBroker()
	go b.Start()

	// Listen for incoming tcp connections ie: telnet
	ln, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(*port))
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer func() { _ = ln.Close() }()

	// handle new incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error handling connection: %s", err)
		}

		id := addNewUser("")

		go handleConnection(conn, b, id)
	}
}

// handles each message we (the client) get from the server
func setupClientMsgHandler(id int, b *Broker, conn net.Conn) {
	msgCh := b.Subscribe()
	for {
		msgData := <-msgCh
		msgDetails := msgData.(ChatMessage)

		// don't write message from self, or if author was muted
		if id != msgDetails.id && isMessageAuthorMuted(id, msgDetails.id) == false {
			conn.Write([]byte(msgDetails.msg))
		}
	}
}

func handleConnection(conn net.Conn, b *Broker, id int) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("Client id: %d with username [%s] connected from %s\n", id, getUserName(id), remoteAddr)

	msg := fmt.Sprintf("%s has joined", getUserName(id))
	sendMessage(msg, 0, b)

	printWelcomeMessage(conn, id)

	// subscribe to all messages
	go setupClientMsgHandler(id, b, conn)

	scanner := bufio.NewScanner(conn)

	// handles every message from the client
	for {
		if ok := scanner.Scan(); !ok {
			break
		}

		msg := fmt.Sprintf(scanner.Text())

		// if msg starts with / process as command
		if len(msg) > 0 && msg[0] == '/' {
			handleCommand(msg, conn, id, b)
		} else {
			sendMessage(msg, id, b)
		}
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleCommand(message string, conn net.Conn, id int, b *Broker) {

	if len(message) > 0 && message[0] == '/' {
		switch {

		case message == "/help":

			msg := `> Commands you can run:
/help - See commands available to you
/history - See the last 30 messages
/name - Change your username
/mute [name] - mute another user
/unmute [name] - unmute a user you have muted previously
/quit - disconnect your client` + "\n"
			// msg = formatMessage(msg, 0)
			conn.Write([]byte(msg))

		case message == "/history":
			resp := strings.Join(msgHistory, "")
			resp = "> Recent messages: \n" + resp
			conn.Write([]byte(resp))

		// starts with /name
		case strings.HasPrefix(message, "/name "):
			oldName := getUserName(id)
			newName := strings.Replace(message, "/name ", "", -1)

			nameTaken := hasUserNameBeenTaken(newName)

			if nameTaken == true {
				msg := "> I'm sorry. The userName [" + newName + "] has already been taken. Try another name\n"
				conn.Write([]byte(msg))

				return
			}

			changeUserName(id, newName)

			// send as system message
			msg := fmt.Sprintf("%s has changed their name to %s", oldName, newName)
			sendMessage(msg, 0, b) // send system notification

		case strings.HasPrefix(message, "/mute "):
			userToMute := strings.Replace(message, "/mute ", "", -1)

			muteUser(id, userToMute)

			// success msg
			msg := fmt.Sprintf("%s has muted %s", getUserName(id), userToMute)
			sendMessage(msg, 0, b) // send system notification

		case strings.HasPrefix(message, "/unmute "):
			userToUnMute := strings.Replace(message, "/unmute ", "", -1)

			unMuteUser(id, userToUnMute)

			// success msg
			msg := fmt.Sprintf("%s has UN-muted %s", getUserName(id), userToUnMute)
			sendMessage(msg, 0, b) // send system notification

		case message == "/quit":
			msg := fmt.Sprintf("%s has left", getUserName(id))
			sendMessage(msg, 0, b)
			conn.Write([]byte("> Goodbye\n"))
			conn.Close()

		default:
			conn.Write([]byte("Unrecognized command.\n"))
		}
	}
}

func getTimeStamp() (timestamp string) {
	now := time.Now()

	timestamp = fmt.Sprintf("%d:%d", now.Hour(),
		now.Minute())
	return timestamp
}

func printWelcomeMessage(conn net.Conn, id int) {
	conn.Write([]byte("> Welcome to the chat room. Type /help for a list of available commands.\n> Your username is [" + getUserName(id) + "]\n"))
}

func changeUserName(id int, newName string) {
	state[id].name = newName
	userNameList[newName] = id
}

func hasUserNameBeenTaken(name string) bool {
	if _, ok := userNameList[name]; ok {
		return true
	}
	return false
}

// searches through mutelist of a user. Returns true if msg should be muted
func isMessageAuthorMuted(userIdRequestingMute int, messageAuthorId int) bool {

	muteList := getMuteList(userIdRequestingMute)

	if _, ok := muteList[messageAuthorId]; ok {
		return true
	}

	return false

}

// returns id of newest user added
func addNewUser(userName string) int {
	// newUser := User{name: userName, id: len(state), muteList: map[int]string{}}
	newUser := User{name: userName, id: len(state), muteList: map[int]string{}}
	state = append(state, newUser)

	userNameList[newUser.name] = newUser.id

	return len(state) - 1
}

// look up username
func getUserName(id int) string {
	userName := state[id].name

	if userName != "" {
		return userName
	}

	userName = "Anon" + strconv.Itoa(id)
	return userName
}

// retrieve saved messages in history
func getMsgHistory() string {
	return strings.Join(msgHistory, "")
}

// add messages to history. limits history to 30 messages
func updateMsgHistory(msg string) {
	// remove first item, if message history is 30
	if len(msgHistory) > 29 {
		msgHistory = append(msgHistory[1:], msg)
	} else {
		msgHistory = append(msgHistory, msg)
	}
}

func getIdFromUserName(userName string) int {
	return userNameList[userName]
}

func getMuteList(id int) map[int]string {
	return state[id].muteList
}

// adds list of users to mute
func muteUser(requestorId int, userToMute string) {
	idToMute := getIdFromUserName(userToMute)
	state[requestorId].muteList[idToMute] = "" // hash map for quick and easy access. Don't care about the value
}

func unMuteUser(requestorId int, userToUnMute string) {
	idToUnMute := getIdFromUserName(userToUnMute)
	delete(state[requestorId].muteList, idToUnMute)
}

// adds the metadata to the message
func formatMessage(msg string, id int) string {
	formattedMsg := fmt.Sprintf("%s [%s] %s\n", getTimeStamp(), getUserName(id), msg)
	return formattedMsg
}

// format and send message to all connected clients
func sendMessage(msg string, id int, b *Broker) {
	// avoid phantom message when user disconnects. TODO: probably need to clean up channel
	if msg == "" {
		return
	}
	msg = formatMessage(msg, id)
	msgData := ChatMessage{msg: msg, id: id}
	b.Publish(msgData) // send to all users
}
