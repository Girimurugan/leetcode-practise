package main

import (
	"fmt"
)

/* Sample Data analysis


 */

/*Logic for the brute force

 */

/*Optimal Logic

 */

// Main Function
func canJump(nums []int) bool {

	const BADINDEX = 'B'
	const GOODINDEX = 'G'

	length := len(nums)
	memo := make([]byte, length, length)

	// Create and initialize the tabulation for Bottom up DP
	for i := 0; i < length; i++ {

		memo[i] = BADINDEX

	}

	memo[length-1] = GOODINDEX // this is the trivial state of the last index

	for i := length - 2; i >= 0; i-- {

		furthestJump := min(i+nums[i], length-1)

		for j := i + 1; j <= furthestJump; j++ {
			if memo[j] == GOODINDEX {
				memo[i] = GOODINDEX
				break
			}
		}
	}

	return memo[0] == GOODINDEX

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

	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))

}
