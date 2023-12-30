package main

import "fmt"

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("Adult:", yearsUntilAdult)
	fmt.Println("Drink:", yearsUntilDrinking)
	fmt.Println("Rental:", yearsUntilCarRental)
}

func main() {
	test(4)
	test(10)
	test(22)
}

func getCoords() (x, y int) {
	// x and y are initialized with zero values
	return // automatically returns x and y
}

func altGetCoords() (int, int) {
	var x int
	var y int
	return x, y
}
