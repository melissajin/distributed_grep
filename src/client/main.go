package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"regexp"
	"io/ioutil"
)

type Counter struct {
    mu  sync.Mutex
    x   int
}

var lineCount Counter

func (c *Counter) Add(x int) {
    c.mu.Lock()
    c.x += x
    c.mu.Unlock()
}

func ConnectToServer(command string, address string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Failed to connect to machine %s\n", address)
		return
	}

	defer conn.Close()

	fmt.Fprintf(conn, command + "\n")
	GetResult(conn, address)
}

func GetResult(connection net.Conn, address string) {
	header := "grep results for machine " + address + ":\n"
	grepOut, _ := ioutil.ReadAll(connection)
	str := string(grepOut)
	out := header + str
	fields := strings.Fields(str)
	numLines := fields[len(fields)-1]
	num, _ := strconv.Atoi(numLines)
	lineCount.Add(num)
	
	WriteToFile(out, address)
	fmt.Printf("Lines on machine %s: %s\n", address, numLines)
}

func WriteToFile(grepOut string, address string) {
	machineNum := ExtractMachineNum(address)
	fileHandle, err := os.Create("grep_" + machineNum + "_out.txt")
	if err != nil {
	    log.Printf("Failed create output file for machine %s\n", machineNum)
	    return
	}

	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, grepOut);
	writer.Flush()
}

func ExtractMachineNum(address string) string {
	re, _ := regexp.Compile("-[0-9]+.")
	str := re.FindString(address)
	if str != "" {
		str = str[1:len(str)-1]
	}
	return str
}

func main() {

	args := strings.Join(os.Args[1:], " ")

	var wg sync.WaitGroup
	for i := 1; i < 11; i++ {
		wg.Add(1)

		machineNum := strconv.Itoa(i)
		fileName := "machine." + machineNum + ".log"
		command := "grep " + args + " " + fileName

		if(len(machineNum) == 1) {
			machineNum = "0" + machineNum
		}

		address := "fa17-cs425-g46-" + machineNum + ".cs.illinois.edu:8000"

		go ConnectToServer(command, address, &wg)
	}
	wg.Wait()

	fmt.Printf("Total line count: %d\n", lineCount.x)
}
