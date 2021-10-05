package main

import (
	"fmt"

	"unicode"
)



type NumStack struct{
	items []int
}

// Push Function

func (s *NumStack)push(value int){
	s.items = append(s.items, value)
}

// Pop Function
func (s *NumStack)pop() int{

	length := len(s.items)-1

	if length >= 0{
	popedValued := s.items[length]
	s.items = s.items[:length]
	

	return	popedValued
	}else{
		return -1
	}
}


type StringStack struct{
	items []string
}

// Push Function

func (s *StringStack)push(value string){
	s.items = append(s.items, value)
}

// Pop Function
func (s *StringStack)pop() string{

	length := len(s.items)-1

	if length >= 0{
	popedValued := s.items[length]
	s.items = s.items[:length]
	

	return	popedValued
	}else{
		return ""
	}
}


func decoded(input string)string{

	numStack := &NumStack{}
	stringStack := &StringStack{}

	k := 0
	currentString := ""
	
	// will be using two stack approach

	//"3[a]2[bc]"
	//"3[a2[c]]"
	for _, char := range input{

		// if this is a digit
		if unicode.IsNumber(char) {
			k = int(char)-'0'

		}else if char == '['{
			numStack.push(k)
			stringStack.push(currentString)
			currentString = ""
			k = 0
		}else if char == ']'{

			decodedString := stringStack.pop()
			for i:= numStack.pop(); i > 0 ;i--{
				decodedString = decodedString + currentString
				fmt.Println("i=", i, "decoded=",decodedString)
			}
			currentString = decodedString

		}else{
			currentString = currentString + string(char)
		}
		fmt.Println(k,currentString)
	}

	return currentString
}

func main(){

	testString := "3[a]2[bc]"

	fmt.Println(decoded(testString))

	testString = "3[a2[c]]"

	fmt.Println(decoded(testString))
}