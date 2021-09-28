package main

import "testing"

func TestHappyNumbers(t *testing.T){

	happyNumberCheck := checkHappyNumber(19)

	if happyNumberCheck != true{
		t.Errorf("Happy number check fails, expected %v but got %v","true",happyNumberCheck )
	}else{
		t.Logf("Success")
	}

}