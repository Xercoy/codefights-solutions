package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	testAdd(t)
}

func testAdd(t *testing.T) {
	if add(5, 4) != 9 {
		t.Errorf("Add failure with two positive numbers")
	}

	if add(-5, -8) != -13 {
		t.Errorf("Add failure with two negative numbers")
	}

	if add(5, -9) != -4 {
		t.Errorf("Add failure with one negative and one positive number")
	}
}
