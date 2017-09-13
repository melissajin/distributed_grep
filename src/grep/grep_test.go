package grep

import "testing"

func TestSearchFile(t *testing.T) {
	command := "grep linux test.html"
	SearchFile(command)
} 