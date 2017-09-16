package grep

import (
	"bytes"
	"os/exec"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SearchFile(command string) string {
	cmd := strings.Fields(command)
	args := cmd[1:]

	out, err := exec.Command(cmd[0], args...).Output()
	if err != nil {
		fmt.Printf("Error executing grep: %s\n", err)
	}

	var buffer bytes.Buffer
	buffer.WriteString(string(out))

	re, _ := regexp.Compile("\n")
	numLines := len(re.FindAllString(buffer.String(), -1))
	buffer.WriteString(strconv.Itoa(numLines))
	buffer.WriteString("\n")
	
	return buffer.String()
}