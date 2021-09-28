package main

import (
	"fmt"
)

/* Sample Data analysis
"abcabcbb" - 3
"bbbbb" = 1
"pwwkew" = 3

*/

/*Logic for the brute force

 */

/*Optimal Logic

 */

// Main Function

func longestSubstringLength(input string) int {

	maxLength := 0
	leftIndex := 0

	visitedChars := make(map[byte]int)

	if len(input) == 0 {
		return 0
	}

	for i := 0; i < len(input); i++ {

		if _, exists := visitedChars[input[i]]; exists {

			if visitedChars[input[i]] > leftIndex {
				leftIndex = visitedChars[input[i]]
			}

		}
		if (i - leftIndex + 1) > maxLength {
			maxLength = i - leftIndex + 1
		}
		visitedChars[input[i]] = i+1

	}

	return maxLength

}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(n)

// Space Complexity - O(min(n, M)) - M is the char set

// Driver function
func main() {

	fmt.Println(longestSubstringLength("abcabcbb"))
	fmt.Println(longestSubstringLength("bbb"))
	fmt.Println(longestSubstringLength("pwwkew"))

}
