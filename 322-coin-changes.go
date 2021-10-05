package main

import (
	"fmt"
)

// main function
func coinChange(coins []int, amount int) int {
    
	// Using the bottom up DP

	// Select the DP parameter which changes in the equation
	// The amount is the which changes as we go along

	// create the DP table first
	dp := make([]int,amount+1) // always set the table length one more than the maximum

	// Initialize the DP with the default value
	dp[0] = 0 // initialize the first because there is no need for a coin for 0 amount
	for i:= 1 ; i< amount+1 ; i++{
		dp[i] = amount+ 1 // amount +1 is the max
	}

	// Calculate the DP

	// Intuition
	//dp[m] = Minimum coins required to make m = m - coin

	for m := 1 ; m < amount+1 ; m++{

		minimumChanges := amount+1

		for c := 0 ; c < len(coins) ; c++{

			change := m-coins[c]

			if change < 0{ // can't make it with this coin
				continue
			}
			if dp[change] < minimumChanges{
				minimumChanges = dp[change]
			}

		}
		dp[m] = minimumChanges+1


	}

	if dp[amount] > amount { // this is because the dp is pre-defaulted with the amount +1 if not updated means no coin is possible
		return -1
	}
	return dp[amount]
}


// Driver function
func main(){

	fmt.Println(coinChange([]int{1,2,5},11))
}