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
	"regexp"
	"io/ioutil"
)

type Counter struct {
    mu  sync.Mutex
    x   int
}

/**
 * Atomic add for type Counter
 */
func (c *Counter) Add(x int) {
    c.mu.Lock()
    c.x += x
    c.mu.Unlock()
}

/* This global variable keeps track of the total number of lines outputted
   by grep from all the connected machines*/
var lineCount Counter

/**
 * ConnectToServer - connects to a server, sends over grep command, and gets
 * the output of grep from the server.
 *
 * @param {string} command - grep command to send to the server
 * @param {string} address - address of server to connect to
 * @param {sync.WaitGroup} wg - sync up multiple goroutines
 */
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

/**
 * GetResult - Retrieves the output of grep from the server and
 * updates the lineCount Counter. Writes the output of grep to a
 * file.
 *
 * @param {net.Conn} connection - connection to server
 * @param {string} address - address of server
 */
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

/**
 * WriteToFile - Writes the content to a file named grep_i_out.txt,
 * i corresponding to the machine number. Creates a new file if it
 * does not  exist or overwrites an existing file.
 *
 * @param {string} content - content to write to file
 * @param {string} address - address of server that produced the grep
 * output. Used to create the filename. 
 */
func WriteToFile(content string, address string) {
	machineNum := ExtractMachineNum(address)
	fileHandle, err := os.Create("grep_" + machineNum + "_out.txt")
	if err != nil {
	    log.Printf("Failed create output file for machine %s\n", machineNum)
	    return
	}

	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, content);
	writer.Flush()
}

/**
 * ExtractMachineNum - Extracts the machine number given the address
 * of the server it corresponds to. Addresses will always be in the form:
 * "fa17-cs425-g46-XX.cs.illinois.edu". XX is the machine number to 
 * be extracted.
 *
 * @param {string} address - address of machine
 @ @return {string} number of the machine
 */
func ExtractMachineNum(address string) string {
	re, _ := regexp.Compile("-[0-9]+.")
	str := re.FindString(address)
	if str != "" {
		str = str[1:len(str)-1]
	}
	return str
}

/**
 * Main entrypoint for the client. Constructs a grep command and 
 * spawns 10 goroutines to connect to 10 servers. 
 */
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
