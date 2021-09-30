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

func onePlus(digits []int)[]int{

	results := make([]int,1 ,1)

	for i:= len(digits)-1 ; i>=0 ; i--{

		if digits[i] == 9 {
			// this is the last digit
			digits[i] = 0
		}else{

			digits[i]++
			return digits
		}

		

	}
	// if not returned earlier then the digits are on the 9s. So add a one at the begining
	results[0] = 1
	results = append(results[:1], digits...)
	return	results

}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(N)

// Space Complexity - O(1)

// Driver function
func main(){

	fmt.Println(onePlus([]int{1,2,3}))
	fmt.Println(onePlus([]int{9,9,9}))

}