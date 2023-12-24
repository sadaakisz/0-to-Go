package main

import "fmt"

// In Go, functions receive a copy of the variable

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25

	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar += sendsToAdd
	return sendsSoFar
}

func basic() {
	x := 5
	incrementBugged(x)

	fmt.Println(x)
	// still prints 5,
	// because the increment function
	// received a copy of x

	x = increment(x)
	fmt.Println(x)
}

func incrementBugged(x int) {
	x++
}

func increment(x int) int {
	x++
	return x
}
