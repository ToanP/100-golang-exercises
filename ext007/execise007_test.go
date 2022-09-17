package main

import (
	"testing"
)

func TestEx007(t *testing.T) {

	// check for error
	want := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 2, 2, 4},
		{0, 2, 4, 6, 8},
	}

	got := Ext007(3, 5)

	for row, rowV := range got {
		for col, cellValue := range rowV {
			if want[row][col] != cellValue {
				t.Errorf("Expected at('%v','%v') is '%v' but got '%v'", row, col, want[row][col], cellValue)
			}
		}
	}
}
