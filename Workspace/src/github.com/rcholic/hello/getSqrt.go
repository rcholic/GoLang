package main

import (
"fmt"
"math"
"errors"
)

func MySqrt(f float64) (float64, error) {
	//return error for negative numbers
	if f < 0 {
		return float64(math.NaN()), errors.New("not computable due to negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Print("First example with -1:\n")
	ret1, err1 := MySqrt(-1)

	if err1 != nil {
		fmt.Println("Error! return values are ", ret1, err1)
	} else {
		fmt.Println("It's ok! return values are ", ret1, err1)
	}

	fmt.Print("second example with 5: \n")
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error return values are ", ret2, err2)
	} else {
		
		fmt.Println("it's ok, return : ", ret2, err2)
	}
}
