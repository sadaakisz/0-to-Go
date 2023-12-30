package main

import "fmt"

// Embedded: truck.make
// Nested: truck.car.make

type car struct {
	make  string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of
	// "truck" now also additionally contains all
	// of the field of the car struct
	car
	bedSize int
}

func basic() {
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			make:  "Toyota",
			model: "Camry",
		},
	}
	fmt.Println(lanesTruck.bedSize)

	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)
}

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rate limit:", s.rateLimit)
	fmt.Println("=============================")
}

func main() {
	test(sender{rateLimit: 10000, user: user{name: "Me", number: 13445}})
	test(sender{rateLimit: 5000, user: user{name: "You", number: 23958}})
	test(sender{rateLimit: 2000, user: user{name: "We", number: 35733}})
}
