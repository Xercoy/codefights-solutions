/* Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

Input/Output

[time limit] 4000ms (go)
[input] array.integer inputArray

An array of integers containing at least two elements.

Guaranteed constraints:
2 ≤ inputArray.length ≤ 10,
-1000 ≤ inputArray[i] ≤ 1000.

[output] integer

The largest product of adjacent elements. */
package main

func adjacentElementsProduct(inputArray []int) int {
	largestProduct := 0

	// Iterate the array in pairs.
	for i := 0; i+2 <= len(inputArray); i++ {
		multiplicand := inputArray[i]
		multiplier := inputArray[i+1]
		product := multiplicand * multiplier

		// Base case - set the first computation to be the
		// largest product, having 0 as a default can interfere
		// with the outcome.
		if largestProduct == 0 {
			largestProduct = product
		}

		if product > largestProduct {
			largestProduct = product
		}
	}

	return largestProduct
}
