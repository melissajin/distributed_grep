package server

import (
	"net"
	"fmt"
	"os"
	"strings"
	"io"
)

func Server() {
	ln, err := net.Listen("tcp", "")
	if err != nil {

	}
	for {
		conn, err := ln.Accept()
		if err != nil {

		}
		go SendFile(conn)
		conn.Close()
	}
}

func SendFile(connection net.Conn) {
	var currentByte int64 = 0
	fileBuffer := make([]byte, 1024)

	file, err := os.Open(strings.TrimSpace("machine.i.log"))
	if err != nil {

	}

	for err == nil || err != io.EOF {
		_, err = file.ReadAt(fileBuffer, currentByte)
		currentByte += 1024
		connection.Write(fileBuffer)

		file.Close()
		connection.Close()
	}
	return
}