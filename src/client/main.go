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
	numFlags := flag.NFlag()
	stringPattern := args[numFlags + 1]

	/*bFlag := flag.Bool("b")
	cFlag := flag.Bool("c")

	eFlag := flag.Bool("e")
	fFlag := flag.Bool("f")
	hFlag := flag.Bool("h")
	iFlag := flag.Bool("i")
	lFlag := flag.Bool("l")
	nFlag := flag.Bool("n")
	pFlag := flag.Bool
	vFlag := flag.Bool("v")
	xFlag := flag.Bool("x")
	yFlag := flag.Bool("y")

	flag.Parse()*/

	/*create client socket and request file from other servers*/
	//send node number to server so that it knows which it is?

	for i := 1; i < 11; i++ {
		go ConnectToServer(i)
	}

}

func ConnectToServer(machineNum int) {
	var address bytes.Buffer
	var fileName bytes.Buffer
	conn, err := net.Dial("tcp", "")
	if err != nil {

	}

	logFile := GetFile( , conn)
	grep.SearchFile( , logFile)
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