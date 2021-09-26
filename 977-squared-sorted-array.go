package main

import (
	"fmt"
)

/* Sample Data analysis


 */

/*Logic for the brute force

 */

/*Optimal Logic
find the flip point of the positive b
*/

// Main Function
func squaredSortedArray(input []int)[]int{

	N := len(input)
	result := make([]int, N)
	leftIndex := 0
	rightIndex := N-1
	square := 0

	for i:= N-1;i>=0;i--{

		if abs(input[leftIndex]) > abs(input[rightIndex]){
			square = input[leftIndex]
			leftIndex++
		}else {
			square = input[rightIndex]
			rightIndex--
		}

		//result = append(result,square * square)
		result[i] = square * square
	}

	return result
}

func abs(input int)int{
	if input < 0{
		return input * -1
	}
	return input
}
// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity

// Space Complexity

// Driver function
func main(){

	fmt.Println(squaredSortedArray([]int{-4,-1,0,3,10}))

}