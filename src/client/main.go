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

func ConnectToServer(args string, machineNum string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileName := "machine." + machineNum + ".log"
	command := "grep " + args + " " + fileName

	if(len(machineNum) == 1) {
		machineNum = "0" + machineNum
	}

	address := "fa17-cs425-g46-" + machineNum + ".cs.illinois.edu:8000"
	//address := "localhost:8000"

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Failed to connect to machine %s\n", machineNum)
		return
	}

	defer conn.Close()

	fmt.Fprintf(conn, command + "\n")
	GetResult(conn, machineNum)
}

func GetResult(connection net.Conn, machineNum string) {
	header := "grep results for machine-" + machineNum + ":\n"
	grepOut, _ := bufio.NewReader(connection).ReadString('\xFF')

	out := header + grepOut[:len(grepOut)-1]
	fields := strings.Fields(grepOut)
	numLines := fields[len(fields)-2]
	num, _ := strconv.Atoi(numLines)
	lineCount.Add(num)
	
	WriteToFile(out, machineNum)
	fmt.Printf("Lines on machine %s: %s\n", machineNum, numLines)
}

func WriteToFile(grepOut string, machineNum string) {
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

func main() {

	args := strings.Join(os.Args[1:], " ")

	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go ConnectToServer(args, strconv.Itoa(i), &wg)
	}
	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println(elapsed)
	fmt.Printf("Total line count: %d\n", lineCount.x)
}
