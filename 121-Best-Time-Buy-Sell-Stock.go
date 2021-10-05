package main

import (
	"fmt"
	"math"
)

const INT_MAX = ^uint(0)

func maxProfit(prices []int) int {

	minimumPrice := math.MaxInt64
	maxProfit := 0

	for i := 0; i < len(prices); i++ {

		if prices[i] < minimumPrice {
			minimumPrice = prices[i]
		} else if prices[i]-minimumPrice > maxProfit {
			maxProfit = prices[i] - minimumPrice
		}
		//fmt.Println(minimumPrice)
		//fmt.Println(maxProfit)
	}

	return maxProfit
}

func main() {

	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

}
