package main

import "fmt"

func handleNewConn;
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
