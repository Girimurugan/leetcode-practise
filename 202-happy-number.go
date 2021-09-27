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

// get next

 func getNext(n int)int{

	// 19 -> 1^2 + 9^2 = 82

	sum := 0
	modNum := 0

	for n > 0{
		modNum = n%10
		n = n/10
		sum += modNum * modNum
	}
	return sum
 }

// Main Function
func checkHappyNumber(input int)bool{

	slowRunner := input
	fastRunner := getNext(input)

	for fastRunner !=1 || fastRunner!= slowRunner{

		slowRunner = getNext(slowRunner)
		fastRunner = getNext(getNext(fastRunner))

	}

	return fastRunner == 1

}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity

// Space Complexity

// Driver function
func main(){

	fmt.Println(checkHappyNumber(19))

	//fmt.Println(getNext(19))

}