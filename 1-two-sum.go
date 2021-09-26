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
func twoSum(input []int, sum int)[2]int{

	scannedNumbers := make(map[int]int)
	resultArray:= [2]int{0,0}

	for i:=0; i < len(input) ; i++ {

		if value, exists := scannedNumbers[sum-input[i]]; exists{

			resultArray[0] = i
			resultArray[1] = value
			return resultArray
		}else{
			scannedNumbers[input[i]] = i
		}
	}

	return resultArray
}


// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(N)

// Space Complexity - O(2+N) = O(N)

// Driver function
func main(){

	fmt.Println(twoSum([]int{2,7,11,15},9))

	fmt.Println(twoSum([]int{3,2,4},6))

}