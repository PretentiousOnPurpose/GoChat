package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type user struct {
	userName     string
	ipAddr       string
	downlinkPort int
	uplinkPort   int
}

var activeChatRoom map[string][]user
var numActiveChatRooms int
var numCreations int

func server_init() {
	fmt.Println("---> PretentiousOnPurpose: chat room server")
	fmt.Println("--->\n---> Current Status:")
}

func newChatRoomServer_init() {
	server, err := net.Listen("tcp", ":8193")
	if err != nil {
		log.Fatalln("Some is wrong with the Creation server!")
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("Some is wrong with the connection request at Creation Server!")
		}

		chatRoomID := createNewChatRoom()
		fmt.Fprintf(conn, "%d\n", chatRoomID)
	}
}

func createNewChatRoom() int {
	chatRoomID := strconv.Itoa(int(numCreations))
	activeChatRoom[chatRoomID] = nil
	numActiveChatRooms++
	numCreations++

	return numCreations - 1
}
func closeChatRoom(chatRoomID int) bool {
	chatRoomID_str := strconv.Itoa(chatRoomID)
	_, ok := activeChatRoom[chatRoomID_str]
	delete(activeChatRoom, chatRoomID_str)

	return ok
}

func validChatRoom(chatRoomID_str string) bool {
	_, ok := activeChatRoom[chatRoomID_str]

	return ok
}

func validChatRoomServer_init() {
	server, err := net.Listen("tcp", ":8192")
	if err != nil {
		log.Fatalln("Some is wrong with the Validation server!")
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("Some is wrong with the connection request at Validation Server!")
		}
		reqText, _ := bufio.NewReader(conn).ReadString('\n')
		reqText = strings.TrimSuffix(reqText, "\n")

		if validChatRoom(reqText) {
			fmt.Fprintf(conn, "valid\n")
		} else {
			fmt.Fprintf(conn, "failed\n")
		}
	}

}

func main() {
	go validChatRoomServer_init()
	newChatRoomServer_init()
}
