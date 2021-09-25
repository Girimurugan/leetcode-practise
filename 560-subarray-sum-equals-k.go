package main

import (
	"fmt"
)

/* Sample Data analysis
[1,1,1] => 2

*/

/*Logic for the brute force

Loop on the array twice and keep a tab of all the sums which are equal to k.
This will be order of complexity as O(n^3) and space complexity O(1)

*/

/*Optimal Logic

We will loop into the array. Get the sum position for each index in an Map
If the next numbers sum-k is already in the map then this number is part of the arrat
* then increase the count

*/

// Main Function
func subarraySumEqualK(input []int, k int) int {

	numOfSubArrays := 0
	currentSum := 0
	mapofSumOccurence := make(map[int]int)

	mapofSumOccurence[0] = 1

	for i := 0; i < len(input); i++ {

		currentSum += input[i]

		if value, exists := mapofSumOccurence[currentSum-k]; exists {
			numOfSubArrays += value
			mapofSumOccurence[currentSum]++
		}else{
			mapofSumOccurence[currentSum] = 1
		}
		

	}

	return numOfSubArrays

}



// Time Complexity - O(n)


// Space Complexity - O(n)

// Driver function
func main() {

	fmt.Println(subarraySumEqualK([]int{1,2,3},3))
	fmt.Println(subarraySumEqualK([]int{3,4,7,2,-3,1,4,2},7))

}
