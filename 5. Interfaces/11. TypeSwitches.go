package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("R: Email to %s cost: %.2f\n", address, cost)
		fmt.Println("================")
	case sms:
		fmt.Printf("R: SMS to %s cost: %.2f\n", address, cost)
		fmt.Println("================")
	default:
		fmt.Printf("R: Invalid expense\n")
		fmt.Println("================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "Hello There",
		toAddress:    "a@a.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "SMS",
		toPhoneNumber: "35235235",
	})
	test(invalid{})
}
