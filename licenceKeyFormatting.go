package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {

	// Loop through the string as character from backwards

	length := len(s)
	counter := 0
	formattedKeyinReverse := ""

	for i := length - 1; i >= 0; i-- {

		// ignore the dash character
		if s[i] == '-' {
			continue
		}

		// start a counter for 4 and reset it every time it reaches the 4 counter

		// when it reaches the counter 4 then add the dashes

		if counter == k {
			formattedKeyinReverse = formattedKeyinReverse + "-" + strings.ToUpper(string(s[i]))
			counter = 1
		} else {

			formattedKeyinReverse = formattedKeyinReverse + strings.ToUpper(string(s[i]))
			counter++
		}

	}

	// reverse the string and return
	formattedKey := ""
	for _, v := range formattedKeyinReverse {

		formattedKey = string(v) + formattedKey

	}

	return formattedKey
}

func main(){

	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}
