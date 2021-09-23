package main

import (
	"fmt"
)

// Min Size Subarray sum function
func minSizeSubArraySum(input []int, targetSum int) int {

	// check for entry conditions
	if len(input) == 0  {
		return 0
	}


	// Return value should be MIN, so by default set to Max
	result := len(input)

	// start two pointers starts and end 
	leftIndex := 0
	valueSum := 0
	
	// loop on the input array
	for i:=0; i< len(input); i++{

		valueSum += input[i]

		for valueSum >= targetSum{
	
			result = min(result, i + 1 - leftIndex )
			
			valueSum -= input[leftIndex]
			leftIndex++
		}	
	}

	if result != len(input){
		return result
	}
	return 0
}

func min(a int, b int)int{
if a>b{
	return b
}
return a
}

// Driver function
func main(){

	input:= []int{2,3,1,2,4,3}
	fmt.Println(minSizeSubArraySum(input,7))

}