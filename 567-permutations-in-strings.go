package main

import (
	"fmt"
)

//// SIMILAR to Leet Code 438

/*Logic for the brute force

 */

/*Optimal Logic

fill the map array for both S1 and s2 strings
then start the index from s2 in the s1.length level and keep checking if the array matches

*/

// Main Function
func checkInclusion(sample string, testString string)bool{

	sampleMap := [26]int{}
	testStringMap := [26]int{}


	sampleN := len(sample)
	testStringN := len(testString)

	//entry condition check
	if sampleN > testStringN{
		return false
	}

	for i:=0 ; i < sampleN ; i++{
		sampleMap[sample[i]-'a']++
		testStringMap[testString[i]-'a']++
	}

	for i := 0 ; i < testStringN - sampleN ; i++ {

		if matching(sampleMap,testStringMap) {
			return true
		}

		testStringMap[testString[i+sampleN]-'a']++
		testStringMap[testString[i]-'a']--
		

	}

	return false


}

func matching(map1 [26]int, map2 [26]int) bool{

	for i:=0;i<26;i++{
		if map1[i]!= map2[i]{
			return false
		}
	}
	return true
}
// Base condition if is recursive

// return conditions

// Dry run with tests
 
// Time Complexity - O(l1 + 26* (l2-l1))

// Space Complexity - O(1)

// Driver function
func main(){

	fmt.Println(checkInclusion("ab","eidbaooo"))

}

