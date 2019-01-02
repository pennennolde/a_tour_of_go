package main

import "fmt"

// 型名は変数名の後ろで宣言する
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 5))
}

