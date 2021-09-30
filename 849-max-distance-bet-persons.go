package main

import (
	"fmt"
)

/* Sample Data analysis

Input: seats = [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (i.e. seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
Example 2:

Input: seats = [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
Example 3:

Input: seats = [0,1]
Output: 1

*/

// Main Function
func maxDistToClosest(seats []int) int {

	// [1,0,0,0,1,0,1]

	N := len(seats)

	if N < 1 {
		return 0
	}

	result := 0

	previous := -1
	future := 0

	for i := 0; i < N; i++ {

		leftDistance := 0
		rightDistance := 0

		if seats[i] == 1 {
			previous = i
		} else {

			for future < N && seats[future] == 0 || future < i {
				future++
			}

			if previous == -1 {
				leftDistance = 0
			} else {
				leftDistance = i - previous
			}

			if future == N {
				rightDistance = N
			} else {
				rightDistance = future - i
			}

		}
		result = max(result, min(leftDistance, rightDistance))

	}

	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
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

	fmt.Println(maxDistToClosest([]int{1,0,0,0,1,0,1}))

	

	fmt.Println(maxDistToClosest([]int{1,0,0,0}))

}
