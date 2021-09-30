package main

import (
	"fmt"
)

/* Sample Data analysis
Input: s = "eceba"
Output: 3
Explanation: The substring is "ece" which its length is 3.

Input: s = "ccaabbb"
Output: 5
Explanation: The substring is "aabbb" which its length is 5.

*/

// Main Function

func longestSubStringofTwoDistinctChars(s string) int {

	// Entry conditions
	if len(s) < 3 {
		return len(s)
	}

	maxLength := 0

	leftIndex := 0

	distinctCharsVisited := make(map[byte]int) // the value is the right most index of the distinct charater

	forwardIndex := 0
	for forwardIndex < len(s) {

		distinctCharsVisited[s[forwardIndex]] = forwardIndex
		forwardIndex++

		if len(distinctCharsVisited) == 3 {

			// find the left most character and remove it

			leftmostIndex := len(s)
			for _, value := range distinctCharsVisited {

				leftmostIndex = min(leftmostIndex, value)

			}
			delete(distinctCharsVisited, s[leftmostIndex])
			leftIndex = leftmostIndex + 1
		}

		maxLength = max(maxLength, forwardIndex-leftIndex)

	}

	return maxLength

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity

// Space Complexity

// Driver function
func main() {

	fmt.Println(longestSubStringofTwoDistinctChars("eceba"))
	fmt.Println(longestSubStringofTwoDistinctChars("ccaabbb"))

}
