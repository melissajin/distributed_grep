package client

import (
	"flag"
	"fmt"
	"os"
	"grep"
	"net"
	"io"
	"bytes"
	"strings"
)

func Client() {

	args := os.Args
	flags := args[1:len(args)-1]
	stringPattern := args[len(args)-1]

	command := "grep "
	for i := 0; i < len(flags); i++ {
		command += ("-" + flags[i] + " ")
	}
	command += stringPattern
	/*create client socket and request file from other servers*/
	//send node number to server so that it knows which it is?

	for i := 1; i < 11; i++ {
		go ConnectToServer(command, i)
	}

}

func ConnectToServer(command string, machineNum int) {
	var address bytes.Buffer
	var fileName bytes.Buffer
	conn, err := net.Dial("tcp", "")
	if err != nil {

	}

	logFile := GetFile( , conn)
	grep.SearchFile(command, logFile)
}

func GetFile(fileName string, connection net.Conn) (logFile *os.File) {

	var currentByte int64 = 0

	fileBuffer := make([]byte, 1024)

	file, err := os.Create(strings.TrimSpace(fileName))
	if err != nil {

	}

	for err == nil || err != io.EOF {

		connection.Read(fileBuffer)
		cleanedFileBuffer := bytes.Trim(fileBuffer, "\x00")

		_, err = file.WriteAt(cleanedFileBuffer, currentByte)

		currentByte += 1024

		if err == io.EOF {
			break
		}

	}

	connection.Close()
	file.Close()
	return
}