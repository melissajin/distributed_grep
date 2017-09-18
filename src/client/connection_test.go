package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestConnection(t *testing.T) {

	go StartDummyServer()

	var wg sync.WaitGroup
	wg.Add(1)
	go ConnectToServer("test", "localhost:8000", &wg)
	wg.Wait()

	file, err := os.Open("grep__out.txt")
	if err != nil {
		t.Errorf("Failure to open file")
	}

	b, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Errorf("Failure to read lines from file")
	}
	lines := strings.Split(string(b), "\n")
	if(lines[0] != "grep results for machine localhost:8000:" && lines[1] != "test") {
		fmt.Println(lines[0])
		fmt.Println(lines[1])
		t.Errorf("File doesn't match grep output")
	}

}

func StartDummyServer() {
	l, _ := net.Listen("tcp", ":8000")
	for {
		conn, _ := l.Accept()
		fmt.Fprintf(conn, "test\n" + "\xFF")
		conn.Close()
	}
	l.Close()
}