package main

import (
	"fmt"
	"strconv"
)

/* Sample Data analysis

Input: nums = [0,1,3,50,75], lower = 0, upper = 99
Output: ["2","4->49","51->74","76->99"]
Explanation: The ranges are:
[2,2] --> "2"
[4,49] --> "4->49"
[51,74] --> "51->74"
[76,99] --> "76->99"

Input: nums = [], lower = 1, upper = 1
Output: ["1"]
Explanation: The only missing range is [1,1], which becomes "1".

Input: nums = [], lower = -3, upper = -1
Output: ["-3->-1"]
Explanation: The only missing range is [-3,-1], which becomes "-3->-1".
*/

/*Logic for the brute force

 */

/*Optimal Logic

 */

// Main Function
func missingRanges(nums []int, lower int, upper int) []string {

	missingRanges := []string{}

	previous := lower-1
	
	

	for i:= 0 ; i<=len(nums);i++{
		current := 0 
		

		if i < len(nums){
			current = nums[i]
		}else{
			current = upper+1
		}

		if (previous + 1 <= current - 1) {
			missingRanges = append(missingRanges, formatRange(previous+1, current-1))
		}
		previous = current;
	}

	// currentNumIndex := 0

	// rangeStartNum := 0
	// rangeEndNum := 0

	// for i := lower; i <= upper; i++ {

	// 	if i == nums[currentNumIndex] {
	// 		if currentNumIndex < len(nums)-1 {
	// 			currentNumIndex++
	// 		}

	// 		if rangeStartNum != 0 && rangeEndNum != 0 {

	// 			missingRanges = append(missingRanges, formatRange(rangeStartNum, rangeEndNum))
			
	// 			rangeStartNum = 0
	// 			rangeEndNum = 0

	// 		}

	// 	} else {

	// 		if rangeStartNum == 0 {
	// 			rangeStartNum = i
	// 		}
	// 		rangeEndNum = i

	// 	}

	// }

	return missingRanges

}

func formatRange(lower int, upper int)string{
	if lower == upper{
		return strconv.Itoa(lower)
	}else{
		return strconv.Itoa(lower) + "->" + strconv.Itoa(upper)
	}
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity

// Space Complexity

// Driver function
func main() {

	fmt.Println(missingRanges([]int{0, 1, 3, 50, 75}, 0, 99))

	//fmt.Println(formatRange(2,2))
	//fmt.Println(formatRange(2,30))

}
