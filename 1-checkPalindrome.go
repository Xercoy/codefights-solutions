/* Given the string, check if it is a palindrome.

Example

For inputString = "aabaa", the output should be
checkPalindrome(inputString) = true;
For inputString = "abac", the output should be
checkPalindrome(inputString) = false;
For inputString = "a", the output should be
checkPalindrome(inputString) = true.
Input/Output

[time limit] 4000ms (go)
[input] string inputString

A non-empty string consisting of lowercase characters.

Guaranteed constraints:
1 ≤ inputString.length ≤ 10.

[output] boolean

true if inputString is a palindrome, false otherwise. */

package main

func checkPalindrome(inputString string) bool {
	// Start indexes
	i := 0
	j := len(inputString) - 1

	// Start from the left and right, comparing strings until
	// there's no match. The break condition simplifies checking
	// for even and odd strings.
	for true {
		leftIndex := i
		rightIndex := j

		if leftIndex >= rightIndex {
			break
		} else if inputString[leftIndex] != inputString[rightIndex] {
			return false
		}

		i++
		j--
	}

	return true
}
