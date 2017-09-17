package main

import (
	"testing"
	"fmt"
)

func TestExtractNum(t *testing.T) {
	tables := []struct {
		address string
		num string
	} {
		{"adf-93.fdjkalsdjf.df:9103", "93"},
		{"fa17-cs425-g46-01.cs.illinois.edu:8000", "01"},
		{"-11.", "11"},
	}

	for _, table := range tables {
		num := ExtractMachineNum(table.address)
		if num != table.num {
			fmt.Println(num)
			t.Errorf("Incorrect machine number extracted")
		}
	}
}
