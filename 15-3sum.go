package main

import (
	"fmt"
	"sort"
)

/* Sample Data analysis


 */

/*Logic for the brute force

 */

/*Optimal Logic

 */

// Main Function

func sumTriplets(input []int)[][]int{

result := [][]int{}

// first sort the input
sort.Ints(input)

// since we need unique triplet we can ignore the next duplicates

for i:=0; i< len(input) && input[i] <= 0 ; i++ {
	if i==0 || input[i-1] != input[i]{
		getTriplets(input, i, &result)
	}
}
//fmt.Println(result)
return result
}

func getTriplets(input []int, index int, result *[][]int){
	leftIndex := index + 1
	rightIndex := len(input)-1
	sum := 0
	
	for leftIndex < rightIndex{
		sum = input[index]+input[leftIndex]+ input[rightIndex]

		triplet := []int{}
		if sum < 0{
			leftIndex++
		}else if sum > 0{
			rightIndex--
		}else {
			triplet = append(triplet, input[index])
			triplet = append(triplet, input[leftIndex])
			leftIndex++
			triplet = append(triplet, input[rightIndex])
			rightIndex--
			//fmt.Println(triplet)
			*result = append(*result, triplet)
			
			//fmt.Println(*result)
		
			for leftIndex<rightIndex && input[leftIndex-1] == input[leftIndex]{
				leftIndex++
			}
		}	

	}

}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity = O (nlogn + n^2) = O(n^2)

// Space Complexity = O(n)

// Driver function
func main(){

	input := []int{-1,0,1,2,-1,-4}

	fmt.Println(sumTriplets(input))

}