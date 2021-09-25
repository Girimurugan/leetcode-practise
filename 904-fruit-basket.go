package main

import (
	"fmt"
)

//// SIMILAR to Leet Code 438
/* Sample Data analysis


 */

/*Logic for the brute force

 */

/*Optimal Logic
Using the Hap map and keep the count with in 2 in the hash map
*/

// Main Function
func maxFruitsInBasket(input []int) int{

	fruitsInBaskets := make(map[int]int)

	leftIndex := 0
	resultMaxFruits := 0

	for rightIndex := 0; rightIndex <len(input) ; rightIndex++{

		if _,exists := fruitsInBaskets[input[rightIndex]] ; exists{
			fruitsInBaskets[input[rightIndex]]++
		}else{
			fruitsInBaskets[input[rightIndex]] = 1
		}

		for len(fruitsInBaskets) > 2{

			fruitsInBaskets[input[leftIndex]]--
			if fruitsInBaskets[input[leftIndex]] == 0{
				delete(fruitsInBaskets,input[leftIndex])
			}
			leftIndex++
		}

		if resultMaxFruits < (rightIndex + 1 - leftIndex){
			resultMaxFruits = rightIndex + 1 - leftIndex
		}

	}

	return resultMaxFruits
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(n)

// Space Complexity - O (1)

// Driver function
func main(){

	fmt.Println(maxFruitsInBasket([]int{1,2,3,2,2}))
	fmt.Println(maxFruitsInBasket([]int{3,3,3,1,2,1,1,2,3,3,4}))

	

}