package main

import (
	"os"
	"net"
	"strings"
	"log"
	"fmt"
	"bufio"
	"sync"
	"strconv"
)

func ConnectToServer(command string, machineNum string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileName := "machine." + machineNum + ".log"
	command = command + " " + fileName

	if(len(machineNum) == 1) {
		machineNum = "0" + machineNum
	}

	address := "fa17-cs425-g46-" + machineNum + ".cs.illinois.edu:8000"

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Failed to connect to machine %d\n", machineNum)
		return
	}

	fmt.Fprintf(conn, command + "\n")
	GetResult(conn, machineNum)
}

func GetResult(connection net.Conn, machineNum string) {
	header := "grep results for machine-" + machineNum + ":\n"
	grepOut, _ := bufio.NewReader(connection).ReadString('\xFF')

	out := header + grepOut[:len(grepOut)-1]
	fmt.Print(out)

	connection.Close()
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
	for i := 1; i < 11; i++ {
		wg.Add(1)
		go ConnectToServer(command, strconv.Itoa(i), &wg)
	}

	wg.Wait()
}
