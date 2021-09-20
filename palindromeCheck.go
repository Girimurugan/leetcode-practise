package main

import (
	"fmt"
)

func checkPalindrome(input string )bool{

	// get the start and end index

	// loop from start and end until, they don't over all
	for startIndex, endIndex := 0,len(input)-1; startIndex<endIndex ; startIndex, endIndex = startIndex+1, endIndex-1{

		if input[startIndex] != input[endIndex]{
			return false
		}

	}

	return true
}

func main(){

	
	fmt.Println(checkPalindrome("india"))
	fmt.Println(checkPalindrome("noon"))
	fmt.Println(checkPalindrome("noaon"))
	fmt.Println(checkPalindrome("nn"))
	fmt.Println(checkPalindrome("b"))
}