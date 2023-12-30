package main

import "fmt"

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

func basic() {
	myCar := car{}
	myCar.FrontWheel.Radius = 5
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func test(m messageToSend) {
	if canSendMessage(m) {
		fmt.Printf("Sending '%s' from (%v) to %s", m.message, m.sender.number, m.recipient.name)
	}
}

func main() {
	test(
		messageToSend{
			message:   "you have an event tomorrow",
			sender:    user{name: "Me", number: 258365},
			recipient: user{name: "Suzie Hall", number: 214847},
		})
	test(
		messageToSend{
			message:   "you have an event tomorrow",
			sender:    user{name: "Me", number: 258365},
			recipient: user{name: "Suzie Hall"},
		})
}
