// func init_client
// func connect to server with chat room ID or create new room req.
// func to chat
// cmd processing to exit or move to another chat room or start over.

package main

import "fmt"

func init_func() {
	fmt.Println("---> Welcome to PretentiousOnPurpose")
	fmt.Println("-----> Premium CHAT ROOM")
	fmt.Println("----->")
	fmt.Println("-------> What would you like to do?")
	fmt.Println("---------> 1. Start a new chat room")
	fmt.Println("---------> 2. Join an existing chat room")
}

func requestUserInput() int {
	var usrInput int
	var userName string
	var chatRoomID int

	fmt.Print("---------> User: ")
	fmt.Scanln(&usrInput)
	fmt.Print("---------> Enter your username: ")
	fmt.Scanln(&userName)

	if usrInput == 1 {
		// createNewChatRoom(&chatRoomID)
		chatRoomID = 100001
	} else {
		fmt.Print("---------> Enter the chat room ID: ")
		fmt.Scanln(&chatRoomID)
		// joinChatRoom(chatRoomID)
	}

	fmt.Println("------->")
	fmt.Printf("-------> Tranferring %s to chat room ID: %d...\n", userName, chatRoomID)

	return 1
}

func connectToChatRoom(chatRoomID int)

func main() {
	init_func()
	chatRoomID := requestUserInput()
	// portID := connectToChatRoom(chatRoomID)

	// open new terminal connecting to the port ID and show chat
	// current terminal acts as keyboard input for sending msgs or other cmds

}
