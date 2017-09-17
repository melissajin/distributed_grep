package main

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestWriteToFile(t *testing.T) {
	tables := []struct {
		grepOut string
		address string
		fileName string
	} {
		{"test", "-1234.", "grep_1234_out.txt"},
	}

	for _, table := range tables {
		WriteToFile(table.grepOut, table.address)
		file, err := os.Open(table.fileName)
		if err != nil {
			t.Errorf("Unable to open file")
		}

		b, err := ioutil.ReadFile(file.Name())
		if err != nil {
			t.Errorf("Error reading file")
		}
		content:= string(b)

		if content != (table.grepOut + "\n") {
			t.Errorf("File written incorrectly")
		}
	}
}