package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	testAdd(t)
	testCenturyFromYear(t)

}

func testCheckPalindrome(t *testing.T) {
	inputOutputMap := map[string]bool{
		"aabaa": true,
		"abac":  false,
		"a":     true,
		"az":    false,
	}

	for input, output := range inputOutputMap {
		actualResult := checkPalindrome(input)

		if actualResult != output {
			t.Errorf("Failure calculating palindrome, result yielded %v, should have been %v", actualResult, output)
		}
	}
}

func testCenturyFromYear(t *testing.T) {
	inputOutputMap := map[int]int{
		1905: 20,
		1700: 17,
		1988: 20,
		2000: 20,
		2001: 21,
		200:  2,
		374:  4,
	}

	for input, output := range inputOutputMap {
		actualResult := centuryFromYear(input)

		if actualResult != output {
			t.Errorf("Failure calculating century for year %d. Caclulated %d, should be %d", input, actualResult, output)
		}
	}
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
