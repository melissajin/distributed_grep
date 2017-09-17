package grep

import (
	"testing"
	"strconv"
)

func TestGrep(t *testing.T) {
	tables := []struct {
		command string
		count int
	} {
		{"grep -c LUKE test.log", 101936},
		{"grep -c force test.log", 2384},
	}

	for _, table := range tables {
		grepOut := SearchFile(table.command)
		count, err := strconv.Atoi(grepOut[:len(grepOut) - len("\n1\n")])
		if err != nil {
			t.Errorf("Failed to retrieve line count")
		}
		if count != table.count {
			t.Errorf("Incorrect number of lines")
		}
	}
}
