package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("Lane,", " happy birthday!"))
	fmt.Println(concat("Elon,", " hope that Tesla thing"))
	fmt.Println(concat("Go", " is fantastic"))
}

func sub(x int, y int) int {
	return x - y
}
