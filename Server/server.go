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

func handleNewConn;

func requestUserInput() int {
	var usrInput int
	var chatRoomID int

	fmt.Print("---------> User: ")
	fmt.Scanf("%d", &usrInput)

	if usrInput == 1 {
		createNewChatRoom(&chatRoomID)
	} else {
		fmt.Println("---------> Enter the chat room ID: ")
		joinChatRoom(chatRoomID)
	}
}

func createNewChatRoom()
func joinChatRoom()

func server_init() {
	fmt.Println("---> PretentiousOnPurpose: chat room server")
	fmt.Println("--->\n---> Current Status:")
	fmt.Println("-----> Number of chat rooms: " + getNumChatRooms())
	// fmt.Println("-----> Number of chat rooms: " + getNumChatRooms())
}

func main() {
	// Start server.
	server_init()
	// [print func]: List currently active chat rooms - their IDs and their initiating usernames
	// start with init_func for every new user conn.

	// init_func()
	// usrInput := requestUserInput()
}
