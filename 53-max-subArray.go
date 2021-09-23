package main

import (
	"fmt"
)

// function to find the max sum of sub array

func maxSumOfSubArray(input []int) int{

	currentSubArray := input[0]
	maxSubArray := input[0]

	// loop through the array

	for i := 1 ; i <len(input) ; i++{

		currentSubArray = max(input[i], currentSubArray + input[i])
		maxSubArray = max(maxSubArray, currentSubArray)


	}

	// if the sum is maximum than the previous result add, else move the pointer



return 	maxSubArray 

}

func max(a int, b int)int{
	if a>b{
		return a
	}
	return b
}

// Driver function
func main(){

	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	input2 := []int{5,4,-1,7,8}
	input3 := []int{1}

	fmt.Println(maxSumOfSubArray(input))
	fmt.Println(maxSumOfSubArray(input2))
	fmt.Println(maxSumOfSubArray(input3))
   
}