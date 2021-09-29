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

// Main Function
func rotateMatrix(matrix [][]int) {

	transpose(matrix)
	reflect(matrix)

}

func transpose(matrix [][]int) {

	length := len(matrix)

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}

func reflect(matrix [][]int) {

	length := len(matrix)

	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(N^2)

// Space Complexity - O(N^2)

// Driver function
func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matrix)

	rotateMatrix(matrix)

	fmt.Println(matrix)

}
