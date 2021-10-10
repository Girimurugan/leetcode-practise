package main

import (
	"fmt"
	"math/rand"
)



//Quickselect (Hoare's selection algorithm)
func topKFrequentQuickSelect(nums []int, k int) []int {

	// Find the frequency count
	countMap := make(map[int]int)

	for _, num := range nums {

		if _, exists := countMap[num]; !exists {
			countMap[num] = 0
		}
		countMap[num]++
	}
    
	// Find the unique numbers array

	N := len(countMap)
	var uniqueNums = make([]int,N)
	i := 0

	for key, _ := range countMap{
		uniqueNums[i] = key
		i++
	}

	fmt.Println(countMap)
	fmt.Println(uniqueNums)

	QuickSelect(&uniqueNums, countMap, 0, N-1, N-k)

	return uniqueNums[N-k:N]

}

func QuickSelect(uniqueNums *[]int, countMap map[int]int, left int, right int, ksmallest int){

	// Base condition
	if left == right{
		return
	}

	pivotIndex:= left+ rand.Intn(right-left)

	pivotIndex = partition(uniqueNums, countMap, left, right, pivotIndex)

	if pivotIndex == ksmallest{
		return
	}else if ksmallest < pivotIndex{
		// go left
		QuickSelect(uniqueNums, countMap, left, pivotIndex-1, ksmallest)
	}else{
		// go right
		QuickSelect(uniqueNums, countMap, pivotIndex+1, right, ksmallest)
	}
}

func partition(uniqueNums *[]int, countMap map[int]int, left int, right int, pivotIndex int)int{
	pivotFrequency := countMap[pivotIndex]

	(*uniqueNums)[pivotIndex], (*uniqueNums)[right] = (*uniqueNums)[right], (*uniqueNums)[pivotIndex]

	storeIndex :=  left

	for i:=left; i <=right ; i++{
		if countMap[i] < pivotFrequency{
			(*uniqueNums)[storeIndex], (*uniqueNums)[i] = (*uniqueNums)[i], (*uniqueNums)[storeIndex]
			storeIndex++
		}
		fmt.Println(uniqueNums)
	}

	(*uniqueNums)[storeIndex], (*uniqueNums)[right] = (*uniqueNums)[right], (*uniqueNums)[storeIndex]
	return storeIndex

}

func main() {

	fmt.Println(topKFrequentQuickSelect([]int{1, 1, 1, 2, 2, 3}, 2)) // [1,2]
}
