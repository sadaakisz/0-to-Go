package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

func basics() {
	height := 10

	if height > 4 {
		fmt.Println("You are tall enough!")
	}

	if height > 6 {
		fmt.Println("Super tall")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
	}
}
