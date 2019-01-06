package main

import "fmt"

func main() {
	sum := 1
	// for文の初期化と後処理ステートメントの記述は任意、; は書く
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

