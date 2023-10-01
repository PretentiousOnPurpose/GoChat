package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func server_init() {
	fmt.Println("---> PretentiousOnPurpose: chat room server")
	fmt.Println("--->\n---> Current Status:")
	// fmt.Println("-----> Number of chat rooms: " + getNumChatRooms())
	// fmt.Println("-----> Number of chat rooms: " + getNumChatRooms())
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

		if reqText == "100001\n" {
			fmt.Fprintf(conn, "valid\n")
		} else {
			fmt.Fprintf(conn, "failed\n")
		}
	}

}

func main() {
	// Start server.
	// server_init()
	validChatRoomServer_init()
	// [print func]: List currently active chat rooms - their IDs and their initiating usernames
	// start with init_func for every new user conn.

	// init_func()
	// usrInput := requestUserInput()
}
