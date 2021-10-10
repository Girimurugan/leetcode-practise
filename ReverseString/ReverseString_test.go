package main

import "testing"



func TestReverseString(t *testing.T){

	reversedString := reverseString("India")

	if reversedString != "aidnI"{
		t.Errorf("String is not reverse properly")
	}
}