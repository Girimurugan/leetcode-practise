package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	const MaxUint = ^uint(0)
	const MinUint = int(0)


	
	fmt.Println(MaxUint)
	fmt.Println(MaxUint > 9999999)
	fmt.Println(MinUint)

	value := int(-3432424242)

	fmt.Println(math.Inf(-1) < float64(value))

	fmt.Println(int(math.Inf(-1)))
	fmt.Println(uint(math.Inf(1)))

	const MAXINT = ^uint(0)
	const MININT = ^int(0)

	fmt.Println((MININT))
	fmt.Println((MAXINT))

	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	fmt.Println(strconv.Atoi(string('1')))

	

	
}