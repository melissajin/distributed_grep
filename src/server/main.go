package main

import (
	"net"
	"log"
	"fmt"
	//"io/ioutil"
	//"os"
	//"strings"
	//"io"
	"os"
	//"strings"
	//"io"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println("Failed to connect to client\n")
	}
	fmt.Println("Listening to port 8000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept client connection\n")
		}
		SendFile(conn)
		conn.Close()
		break
	}
}

func SendFile(connection net.Conn) {
	defer connection.Close()

	file, err := os.Open("machine.1.log")
	if err != nil {
		log.Printf("Failed to open file %s\n", "machine.1.log")
	}
	defer file.Close()
	fi,_ := file.Stat()
	size := fi.Size()

	fileBuffer := make([]byte, size)

	n, err := file.ReadAt(fileBuffer, 0)
	connection.Write(fileBuffer[:n])

	return
}