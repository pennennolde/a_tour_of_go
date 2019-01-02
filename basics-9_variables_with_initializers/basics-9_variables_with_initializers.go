package main

import "fmt"

// var宣言では初期化子を与えることができる
// 初期化する場合は型宣言を省略できる
var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

