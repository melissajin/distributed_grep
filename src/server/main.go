package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"grep"
)

func main() {

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("Failed to connect to client\n")
	}
	fmt.Println("Listening to port 8000")

	for {
		conn, err := ln.Accept()
		fmt.Println("Accepted connection")
		if err != nil {
			log.Println("Failed to accept client connection\n")
		}
		command, _ := bufio.NewReader(conn).ReadString('\n')
		grepOut := grep.SearchFile(command)
		SendOutput(conn, grepOut)
		conn.Close()
	}
}

func SendOutput(connection net.Conn, grepOut []byte) {
	out := string(grepOut)
	fmt.Fprintf(connection, out + "\xFF")
}