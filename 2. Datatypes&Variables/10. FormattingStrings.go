package main

import "fmt"

func main() {
	// fmt.Printf allows the following interpolations:
	// %v: default representation
	// %s: interpolate string
	// %d: interpolate an int in decimal form
	// %f: interpolate a decimal
	// %.2f: interpolate a decimal with 2 decimal places

	// fmt.Sprintf returns the formatted string and doesn't print to console

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)
}
