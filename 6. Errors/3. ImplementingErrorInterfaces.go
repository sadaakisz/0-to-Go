package main

import "fmt"

type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their acc", e.name)
}

func sendSMS(msg, userName string) error {
	return userError{name: userName}
}

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", e.dividend)
}

func divide(divident, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideError{dividend: divident}
	}
	return divident / divisor, nil
}

func test(divident, divisor float64) {
	defer fmt.Println("=======================")
	fmt.Printf("Dividing %.2f by %.2f ... \n", divident, divisor)
	quotient, err := divide(divident, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
