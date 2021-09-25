package main

import (
	"fmt"
)

/* Sample Data analysis


 */

/*Logic for the brute force


 */

/*Optimal Logic
Create an Map whose eize does not cross K+1
SLiding windoe concept

*/

// Main Function
func longestSubstringLenKDistChars(input string, k int)int{

	lengthOfInput := len(input)
	if lengthOfInput == 0 || k == 0{
		return 0
	}

	maxLength := 0

	leftIndex := 0
	rightIndex := 0

	charMaps := make(map[byte]int)

	for rightIndex < lengthOfInput{

		/*if value,exists := charMaps[input[rightIndex]];exists{
			charMaps[input[rightIndex]] = rightIndex
		}*/

		charMaps[input[rightIndex]] = rightIndex
		rightIndex++

		// if the maximum k chars is reached, remove the left move char
		if len(charMaps) == k + 1 {

			leftMostChar := getleftMostChar(charMaps)
			leftIndex = charMaps[leftMostChar] + 1
			delete(charMaps,leftMostChar)
			

		}
		if rightIndex - leftIndex > maxLength{
			maxLength = rightIndex - leftIndex
		}
		

	}


	return maxLength

}

func getleftMostChar(input map[byte]int) byte{

	smallestIndexSofar := 99999
	var result  byte
	for key, value := range input{

		if value < smallestIndexSofar{
			smallestIndexSofar = value
			result = key
		}

	}

	return result
}

// Base condition if is recursive

// Basic input validations

// return conditions

// Dry run with tests

// Time Complexity - O(n)

// Space Complexity - O(k)

// Driver function
func main(){

	fmt.Println(longestSubstringLenKDistChars("eceba",2))

}