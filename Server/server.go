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

func readReqText(conn net.Conn) string {
	data, _ := bufio.NewReader(conn).ReadString('\n')

	return data
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

		reqText := readReqText(conn)

		fmt.Println(reqText)

		if reqText == "100001" {
			fmt.Fprintf(conn, "valid")
		} else {
			fmt.Fprintf(conn, "failed")
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
