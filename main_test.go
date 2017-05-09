package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	testAdd(t)
	testCenturyFromYear(t)
	testCheckPalindrome(t)
	testAdjacentElementProduct(t)
	testShapeArea(t)
}

func testShapeArea(t *testing.T) {
	inputOutputMap := map[int]int{
		2: 5,
		3: 13,
		1: 1,
		5: 41,
	}

	for input, output := range inputOutputMap {
		actualResult := shapeArea(input)

		if actualResult != output {
			t.Errorf("Error calculating n-interesting polygon squares, result yielded %d, should have been %d", actualResult, output)
		}
	}
}
func testAdjacentElementProduct(t *testing.T) {
	inputOutputMap := map[int][]int{
		21:  []int{3, 6, -2, -5, 7, 3},
		2:   []int{-1, -2},
		6:   []int{5, 1, 2, 3, 1, 4},
		50:  []int{9, 5, 10, 2, 24, -1, -48},
		30:  []int{5, 6, -4, 2, 3, 2, -23},
		-12: []int{-23, 4, -3, 8, -12},
	}

	for output, input := range inputOutputMap {
		actualResult := adjacentElementsProduct(input)

		if actualResult != output {
			t.Errorf("Failure calculating largest product, result yielded %d, should have been %d", actualResult, output)
		}
	}
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
