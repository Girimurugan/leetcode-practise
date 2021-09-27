package main

import (
	"fmt"
)

/* Sample Data analysis


 */

/*Logic for the brute force

 */

/*Optimal Logic
Using 3 pointers.. one for current, one for right of the 0s, one for the left of the 2s.
*/

// Main Function
func dutchFlafSorting(input []int)[]int{

	
	zerosRightendIndex := 0
	twosleftendIndex := len(input)-1
	currentIndex:=0

	for currentIndex<=twosleftendIndex{

		if input[currentIndex] == 0{
			input[currentIndex], input[zerosRightendIndex] = input[zerosRightendIndex], input[currentIndex]
			zerosRightendIndex++
			currentIndex++

		}else if input[currentIndex] == 2{
			input[currentIndex], input[twosleftendIndex] = input[twosleftendIndex], input[currentIndex]
			twosleftendIndex--
		}else{
			currentIndex++
		}

	}


	return input

}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(n)

// Space Complexity - O(1)

// Driver function
func main(){

	fmt.Println(dutchFlafSorting([]int{2,0,2,2,1,1,0}))

}