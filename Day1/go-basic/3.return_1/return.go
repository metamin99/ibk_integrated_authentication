package main

import "fmt"

func main() {
	fmt.Println(testReturn())
}

func testReturn() (string, string, int) {
	return "3", "1", 1
}
