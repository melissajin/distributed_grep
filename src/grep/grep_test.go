package grep

import "testing"

func TestSearchFile(t *testing.T) {
	command := "grep linux index.html"
	SearchFile(command)
} 