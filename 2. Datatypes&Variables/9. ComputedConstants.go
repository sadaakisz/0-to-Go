package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	// only valid because secondsInMinute and minutesInHour are know values
	// this is computed in compile time
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("Number of seconds in an hour:", secondsInHour)
}
