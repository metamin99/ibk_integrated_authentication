package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now().Format("20060102150405")  // 20060102150405 은 고정된 숫자
	// yyyy-mm-dd-hh-mm-ss 의미하여 실제로 2006-01-02-15-04-05 로 넣어도 됨
	p(now)
}
