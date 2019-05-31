package main

import "fmt"

func main() {
	var1, _ := testReturn()
	fmt.Println(var1)
	//fmt.Println(var2)
}

func testReturn() (string, string) {
	return "3", "2"
}
