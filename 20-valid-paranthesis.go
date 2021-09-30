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

type stack struct{

	items []byte

}

func (s *stack)push(bracket byte){
	
	s.items = append(s.items, bracket)
	
}

func (s * stack)pop()byte{

	toPop := byte(0)
    if len(s.items) > 0{
	toPop = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
    }

	return toPop
}

// Main Function
func checkValidParanthesis(s string) bool{

	bracketStack := &stack{}

	
	charMaps :=  make(map[byte]byte)
	charMaps[')'] = '('
	charMaps['}'] = '{'
	charMaps[']'] = '['


	for i:= 0 ; i< len(s);i++{

		if value, exists := charMaps[s[i]]; exists{

			topOfStack := bracketStack.pop()
			if  topOfStack != value{
				return false
			}
		}else{
			bracketStack.push(s[i])
		}

	}

	return len(bracketStack.items) == 0

}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity - O(n)

// Space Complexity - O(n)

// Driver function
func main(){

	fmt.Println(checkValidParanthesis("()[]{}"))

	fmt.Println(checkValidParanthesis("{[]}"))

	fmt.Println(checkValidParanthesis("([)]"))

	

}