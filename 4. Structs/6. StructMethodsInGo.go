package main

import "fmt"

type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}

func basic() {
	r := rect{
		width:  5,
		height: 10,
	}
	fmt.Println(r.area())
}

type authInfo struct {
	username string
	password string
}

func (a authInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func test(aI authInfo) {
	fmt.Println(aI.getBasicAuth())
	fmt.Println("========================")
}

func main() {
	test(authInfo{username: "Me", password: "23234"})
	test(authInfo{username: "U", password: "32458"})
	test(authInfo{username: "We", password: "56938"})
}
