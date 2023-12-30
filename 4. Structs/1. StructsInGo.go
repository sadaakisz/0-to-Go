package main

import "fmt"

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("============================================")
}

func main() {
	test(messageToSend{phoneNumber: 148255510981, message: "Thanks for signing up"})
	test(messageToSend{phoneNumber: 325646463466, message: "Thx"})
}
