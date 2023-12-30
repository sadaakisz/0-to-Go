package main

import (
	"errors"
	"fmt"
)

func main() {
	x, _ := divide(4, 2)
	y, _ := divide(2, 0)
	fmt.Println(x, y)
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// Guard clause
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}
