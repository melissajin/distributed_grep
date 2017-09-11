package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"grep"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println("Failed to connect to client\n")
	}
	fmt.Println("Listening to port 8000")

	conn, err := ln.Accept()
	if err != nil {
		log.Println("Failed to accept client connection\n")
	}

	for {
		command, _ := bufio.NewReader(conn).ReadString('\n')
		grepOut := grep.SearchFile(command)
		SendOutput(conn, grepOut)
		conn.Close()
		break
	}
}

func SendOutput(connection net.Conn, grepOut []byte) {
	out := string(grepOut)
	fmt.Fprintf(connection, out + "\xFF")
}