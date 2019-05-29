package main

import "fmt"

func main() {
	var a = "this is a"
	fmt.Println(a)  // go 언어는 대문자로 지정하면 퍼블릭, 소문자는 로컬(퍼블릭 접근 안됨)

	b := "this is b"  // := 는 변수 초기화 의미, 타입은 지정된 양식에 따라 변수가 자동으로 형지정(변환)됨
	b = "this is c"
	fmt.Println(b)
}
