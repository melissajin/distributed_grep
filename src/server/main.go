package main

import (
	"bufio"
	"fmt"
	"grep"
	"log"
	"net"
)

/**
 * Main entrypoint for the server. Creates a new server
 * listening to port 8000 and waits to accept connections
 * from a client. Once accepted, it will execute the grep
 * command sent from the client and send the output of grep
 * back to the client.
 */
func main() {

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("Failed to connect to client\n")
	}
	fmt.Println("Listening to port 8000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept client connection\n")
			continue
		}

		fmt.Println("Accepted connection")
		command, _ := bufio.NewReader(conn).ReadString('\n')
		grepOut := grep.Execute(command)
		conn.Write([]byte(grepOut))
		conn.Close()
	}
}