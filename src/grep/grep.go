package grep

import (
	 "os/exec"
	"fmt"
	"strings"
)

func SearchFile(command string) {
	cmd := strings.Fields(command)
	args := cmd[1:]

	out, err := exec.Command(cmd[0], args...).Output()

	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	
	fmt.Printf("%s", out)
}