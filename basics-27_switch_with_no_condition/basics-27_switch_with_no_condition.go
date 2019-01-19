package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	h := t.Hour() + 9
	fmt.Println(h)
	// 長くなりがちな"if-then-else"をシンプルに表現できる
	switch {
	// switch true { と同じ
	case h < 12:
		fmt.Println("Good morning!")
	case h < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

