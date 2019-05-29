package main

import "fmt"

func main() {
	a, b := testReturn()
	fmt.Println(a)
	fmt.Println(b)
}

func testReturn() (string, string) {
	return "3", "1"
}
