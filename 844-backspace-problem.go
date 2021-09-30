package main

import (
	"fmt"
)



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
func (s *stack) getContentAsString()string{

	result := ""

	for i:=0; i< len(s.items); i++{
		result = result + string(s.items[i])
	}
	return result
}

func backspaceCompare( s string, t string)bool{

	if applyBackspace(s) == applyBackspace(t){
		return true
	}

	return false
}

func applyBackspace(s string)string{

	charStack := &stack{}

	for i:= 0; i< len(s) ; i++{
		if s[i] != '#'{
			charStack.push(s[i])
		}else{
			charStack.pop()
		}
	}

	return charStack.getContentAsString()
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity

// Space Complexity

// Driver function
func main(){

	fmt.Println(backspaceCompare("ab##","c#d#"))

	fmt.Println(backspaceCompare("a##c","#a#c"))

	fmt.Println(backspaceCompare("a#c","b"))

}