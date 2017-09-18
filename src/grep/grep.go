package grep

import (
	"bytes"
	"os/exec"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
 * Execute - executes the grep command by calling the Unix grep.
 * Counts the number of lines outputted by grep and appends the
 * number to the end of the grep output.
 *
 * @param {string} command - grep command to be executed
 * @return {string} output of grep with number of lines appended.
 */
func Execute(command string) string {
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