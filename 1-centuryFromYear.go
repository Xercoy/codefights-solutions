/* Given a year, return the century it is in. The first century spans from the year 1 up to and including the year 100, the second - from the year 101 up to and including the year 200, etc.

Example

For year = 1905, the output should be
centuryFromYear(year) = 20;
For year = 1700, the output should be
centuryFromYear(year) = 17.
Input/Output

[time limit] 4000ms (go)
[input] integer year

A positive integer, designating the year.

Guaranteed constraints:
1 ≤ year ≤ 2005.

[output] integer

The number of the century the year is in. */

/*
Brilliant solution inspired by Tiff:

func centuryFromYear(year int) int {
    return 1 + ((year - 1) / 100)
}

*/

package main

func centuryFromYear(year int) int {
	// The base value for a century is 1.
	century := 1

	// Get the value of the two rightmost digits.
	twoDigitValue := year % 100

	// Get the value of the leftmost digit which is
	// the hundreds place digit.
	hundredsPlaceDigit := year / 100

	// If the two digit value ends in 0, the hundreds
	// place digit already indicates the correct century.
	if twoDigitValue == 0 {
		century = hundredsPlaceDigit
	}

	// If the two digit value is 1 or greater than 1, the
	// hundreds place digit indicates a number that is 1
	// less than the century, so add 1 to it.
	if twoDigitValue >= 1 {
		century = hundredsPlaceDigit + 1
	}

	return century
}
