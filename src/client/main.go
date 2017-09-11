package main

import (
	"os"
	"grep"
	"net"
	"strings"
	"log"
	"fmt"
	"bufio"
	"sync"
	"bytes"
	"io"
)

func ConnectToServer(command string, machineNum int, wg *sync.WaitGroup) {
	defer wg.Done()

	fileName := "machine.1.received.log"
	address := "localhost:8000"

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Failed to connect to machine %d\n", machineNum)
		return
	}

	logFile := GetFile(fileName, conn); if logFile == nil {
		return
	}

	command = command + " " + fileName
	grep.SearchFile(command)
}

func GetFile(fileName string, connection net.Conn) (logFile *os.File) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Failed to create file %s\n", fileName)
		return nil
	}

	var buf bytes.Buffer
	io.Copy(&buf, connection)
	file.WriteAt(buf.Bytes(), 0)

	connection.Close()
	file.Close()
	return file
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	command, _ := reader.ReadString('\n')
	command = strings.Replace(command, "\n", "", -1)

	fields := strings.Fields(command)
	if len(fields) < 2 {
		fmt.Println("Command invalid")
		return
	} else if fields[0] != "grep" {
		fmt.Println("Command must be grep")
		return
	}

	var wg sync.WaitGroup
	for i := 1; i < 2; i++ {
		wg.Add(1)
		go ConnectToServer(command, i, &wg)
	}

	wg.Wait()
}
