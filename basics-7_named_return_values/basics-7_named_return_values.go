package main

import "fmt"

// 戻り値は名前をつけておくことができる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	// "naked" return
}

func main() {
	fmt.Println(split(17))
}

