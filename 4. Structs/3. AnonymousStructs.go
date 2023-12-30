package main

import "fmt"

func main() {
	// Unusual usecase
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Tesla",
		Model: "Model 3",
	}
	fmt.Println(myCar.Make)

	// Usual usecase
	type car struct {
		Make   string
		Model  string
		Height int
		Width  int

		// Wheel is a field containing an anonymous struct
		Wheel struct {
			Radius   int
			Material string
		}
	}
}
