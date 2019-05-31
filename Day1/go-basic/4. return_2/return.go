package main

import "fmt"

func main() {
	a, _ := testReturn()
	fmt.Println(a)
}

func testReturn() (string, string) {
	return "3", "2"
}
