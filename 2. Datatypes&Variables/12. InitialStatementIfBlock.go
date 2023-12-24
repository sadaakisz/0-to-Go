// WON'T RUN, just as a notepad

package main

import "fmt"

func main() {
	length := getLength(email)
	if length < 1 {
		fmt.Println("Email is invalid")
	}

	// SIMILAR TO, length would be accesible in the if scope:

	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid")
	}
}

func basic() {
	if INITIAL_STATEMENT; CONDITION {

	}
}
