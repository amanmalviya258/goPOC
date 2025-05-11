package main

import (
	"errors"
	"fmt"
)

func main() {
	calling()
}

func calling() {

	num1 := 22
	num2 := 7
	num3 := 40.4

	var result, remainder, floatNum, error = calculation(num1, num2, num3)
	if error != nil {
		fmt.Printf(error.Error())
	}

	fmt.Printf("The remainder is %v with result is %v and the num3 is %v", remainder, result, floatNum)
}

func calculation(numerator int, demo int, num3 float64) (int, int, float64, error) {
	var errorhaibhai error //default value is null when it is not initialised similarly how 0 is default value of int.

	if demo == 0 {
		errorhaibhai = errors.New("cannot divide by 0")
		return 0, 0, 0, errorhaibhai
	}
	var result = numerator / demo
	var demomenator = numerator % demo
	return result, demomenator, num3, errorhaibhai
}
