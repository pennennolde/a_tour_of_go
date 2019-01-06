package main

import "fmt"

func main() {
	sum := 0
	// Golangのfor文に()はいらない
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

