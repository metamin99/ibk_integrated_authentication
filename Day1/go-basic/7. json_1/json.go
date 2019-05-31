package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
	RegDate      string `json:"regDate"`
}

func main() {

	//1. create struct data
	user := User{
		UserId:       "user01",
		UserName:     "홍길동",
		UserPassword: "1q2w3e4r@",
		RegDate:      "20190101",
	}

	//2. convert struct to jsonBytes
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	//3. convert jsonBytes to jsonString
	jsonString := string(jsonBytes)

	//4. print jsonString
	fmt.Println(jsonString)
}
