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
func containerWithMostWaterArea(heights []int) int{

maxArea := 0

leftIndex := 0
rightIndex := len(heights)-1

for leftIndex < rightIndex{

	maxArea = max(maxArea, (min(heights[leftIndex], heights[rightIndex])* (rightIndex-leftIndex)))

	if heights[leftIndex] < heights[rightIndex]{
		leftIndex++
	}else{
		rightIndex--
	}

}

return maxArea

}

func max (a int, b int)int{
	if a > b{
		return a
	}
	return b
}


func min (a int, b int)int{
	if a < b{
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
func main(){

	fmt.Println(containerWithMostWaterArea([]int{1,8,6,2,5,4,8,3,7}))

	fmt.Println(containerWithMostWaterArea([]int{4,3,2,1,4}))

	fmt.Println(containerWithMostWaterArea([]int{1,1}))
}