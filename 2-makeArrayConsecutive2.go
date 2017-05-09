/*
Ratiorg got statues of different sizes as a present from CodeMaster for his birthday, each statue having an non-negative integer size. Since he likes to make things perfect, he wants to arrange them from smallest to largest so that each statue will be bigger than the previous one exactly by 1. He may need some additional statues to be able to accomplish that. Help him figure out the minimum number of additional statues needed.

Example

For statues = [6, 2, 3, 8], the output should be
makeArrayConsecutive2(statues) = 3.

Ratiorg needs statues of sizes 4, 5 and 7.

Input/Output

[time limit] 4000ms (go)
[input] array.integer statues

An array of distinct non-negative integers.

Guaranteed constraints:
1 ≤ statues.length ≤ 10,
0 ≤ statues[i] ≤ 20.

[output] integer

The minimal number of statues that need to be added to existing statues such that it contains every integer size from an interval [L, R] (for some L, R) and no other sizes.
*/

/* Genius Answer
func makeArrayConsecutive2(statues []int) int {
    highest := -1
    lowest := 21

    for _, n := range statues {
        if n > highest {
            highest = n
        }
        if n < lowest {
            lowest = n
        }
    }

    return (highest - lowest + 1) - len(statues)
}
*/

package main

func makeArrayConsecutive2(statues []int) int {
	lowestInteger := findLowestInteger(statues)
	highestInteger := findHighestInteger(statues)
	integerDifference := highestInteger - lowestInteger

	return integerDifference - (len(statues) - 1)
}

func findHighestInteger(numList []int) int {
	highestNum := numList[0]

	for i := len(numList) - 1; i > 0; i-- {
		value := numList[i]

		if value > highestNum {
			highestNum = value
		}
	}

	return highestNum
}

func findLowestInteger(numList []int) int {
	lowestNum := numList[0]

	for i := len(numList) - 1; i > 0; i-- {
		value := numList[i]

		if value < lowestNum {
			lowestNum = value
		}
	}

	return lowestNum
}
