package main

import "fmt"

// 関数の引数の型が同じ場合は省略して記述できる
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(105, 45))
}

