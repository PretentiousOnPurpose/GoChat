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
)

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
		// createNewChatRoom(&chatRoomID)
		chatRoomID = 100001
	} else {
		fmt.Print("> Enter the chat room ID: ")
		fmt.Scanln(&chatRoomID)
		// joinChatRoom(chatRoomID)
	}

	fmt.Println("> Processing...")
	// check if valid chatRoomID

	if validChatRoom(chatRoomID) == "valid" {
		fmt.Printf("> Tranferring %s to chat room ID: %d...\n", userName, chatRoomID)
	} else {
		fmt.Printf("> Chat room ID: %d does not exist. Please try again.\n", chatRoomID)
	}

	return 1
}

func validChatRoom(chatRoomID int) string {
	conn, err := net.Dial("tcp", ":8192")
	if err != nil {
		log.Fatalln("Some is wrong with the Validation server!")
	}
	fmt.Println("Here")
	fmt.Fprintf(conn, "%d", chatRoomID)
	// time.Sleep(1)
	fmt.Println("Here2")

	data, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Here3")
	fmt.Println(data)
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
