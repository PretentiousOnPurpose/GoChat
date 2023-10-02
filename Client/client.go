// func init_client
// func connect to server with chat room ID or create new room req.
// func to chat
// cmd processing to exit or move to another chat room or start over.

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

func init_func() {
	fmt.Println("> GoChat chat room")
	fmt.Println("> What would you like to do?")
	fmt.Println("> 1. Start a new chat room")
	fmt.Println("> 2. Join an existing chat room")
}

func requestUserInput() int {
	var usrInput int
	var userName string
	var chatRoomID int

	fmt.Print("> Response : ")
	fmt.Scanln(&usrInput)
	fmt.Print("> Enter your username: ")
	fmt.Scanln(&userName)

	if usrInput == 1 {
		chatRoomID = createNewChatRoom()
	} else {
		fmt.Print("> Enter the chat room ID: ")
		fmt.Scanln(&chatRoomID)
		// joinChatRoom(chatRoomID)
	}

	fmt.Println("> Processing...")
	fmt.Printf("> Tranferring %s to chat room ID: %d...\n", userName, chatRoomID)

	// if validChatRoom(chatRoomID) == "valid\n" {
	// 	fmt.Printf("> Tranferring %s to chat room ID: %d...\n", userName, chatRoomID)
	// } else {
	// 	fmt.Printf("> Chat room ID: %d does not exist. Please try again.\n", chatRoomID)
	// }

	return 1
}

func createNewChatRoom() int {
	conn, err := net.Dial("tcp", ":8193")
	if err != nil {
		log.Fatalln("Some is wrong with the Creation server!")
	}

	chatRoomID_str, _ := bufio.NewReader(conn).ReadString('\n')
	chatRoomID_str = strings.TrimSuffix(chatRoomID_str, "\n")
	chatRoomID, _ := strconv.Atoi(chatRoomID_str)

	return chatRoomID
}

func validChatRoom(chatRoomID int) string {
	conn, err := net.Dial("tcp", ":8192")
	if err != nil {
		log.Fatalln("Some is wrong with the Validation server!")
	}
	fmt.Fprintf(conn, "%d\n", chatRoomID)

	data, _ := bufio.NewReader(conn).ReadString('\n')

	return data
}

func connectChatRoom(chatRoomID int) {

}

func main() {
	init_func()
	_ = requestUserInput()
	// portID := connectToChatRoom(chatRoomID)

	// open new terminal connecting to the port ID and show chat
	// current terminal acts as keyboard input for sending msgs or other cmds

}
