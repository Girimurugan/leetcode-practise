package main

import (
	"fmt"
)

/* Sample Data analysis


 */

/*Logic for the brute force

 */

/*Optimal Logic
Having 2 pointers in the array, one for the place holder for changeable location
Another for the right forward index
*/

// Main Function
func removeDuplicatedInSortedArray(input []int) int {

	leftIndex := 0

	for i := 1; i < len(input); i++ {

		if input[leftIndex] != input[i] {
			leftIndex++
			input[leftIndex] = input[i]
		}

	}

	return leftIndex + 1
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(n)

// Space Complexity - O(1)

// Driver function
func main() {

	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicatedInSortedArray(input)

	fmt.Println(input[:k])


	input = []int{1,1,2}
	k = removeDuplicatedInSortedArray(input)

	fmt.Println(input[:k])

}
