package main

import (
	"fmt"
)

func reverseString(input string)string{

	// base conditions
	if input == ""{
		return ""
	}

	return reverseString(input[1:]) + string(input[0])

}

func main(){

	fmt.Println(reverseString("Giri"))

}