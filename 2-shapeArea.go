/*
Below we will define an n-interesting polygon. Your task is to find the area of a polygon for a given n.

A 1-interesting polygon is just a square with a side of length 1. An n-interesting polygon is obtained by taking the n - 1-interesting polygon and appending 1-interesting polygons to its rim, side by side. You can see the 1-, 2-, 3- and 4-interesting polygons in the picture below.

Example

For n = 2, the output should be
shapeArea(n) = 5;
For n = 3, the output should be
shapeArea(n) = 13.
Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n < 104.

[output] integer

The area of the n-interesting polygon.
*/

/* Mathematical Solution
function shapeArea(n) {
    return n*n + (n-1)*(n-1);
}
*/
package main

func shapeArea(n int) int {
	// Base case when n == 1.
	totalSquares := 1

	// Calculate for all cases where n > 1
	for i := n; i > 1; i-- {
		newSquares := (i - 1) * 4
		totalSquares = totalSquares + newSquares
	}

	return totalSquares
}
