package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

/// Global variables
var PORT string
var server *Server = ConfigureServer()

/// Parses the arguments into their respective variables
func ParseArgs() {
	flag.StringVar(&PORT, "port", "6667", "Port to listen on for connections")
	flag.Parse()

}

func ConfigureServer() *Server {
	server := Server{}
	server.AllowedOpers = make(map[string]string)
	server.AllowedOpers["StatesideCash"] = "youcallthisapassword?"
	return &server
}

/// Handles the connection to the client, as well as maintains their state
func ConnectionHandler(conn net.Conn) {
	fmt.Println("[=]New Connection from:", conn.RemoteAddr())
	defer conn.Close()
	user := User{} /// Stores the users session data
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := ParseMessage(scanner.Text())
		status := HandleCommand(msg, &user, server)
		if status == EXIT_STATUS {
			//TODO Send a message to all the user's chans containing an exit message
			//or a message explaining the network error that caused them to DC
			break
		}
		if status != SUCCESS {
			io.WriteString(conn, string(status)+"\r\n")
		}
		// fmt.Printf("\n%#v\n", msg) /// TODO Remove debug line
		fmt.Printf("%#v\n", user)
	}
	fmt.Println("[=]Stopping Connection:", conn.RemoteAddr())
}

func main() {
	/// Get arguments
	ParseArgs()

	/// Set up the server listener
	server, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	/// Server loop
	for {
		/// Accept incoming connection
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		/// Launch a handler for the new socket
		go ConnectionHandler(conn)
	}
}
