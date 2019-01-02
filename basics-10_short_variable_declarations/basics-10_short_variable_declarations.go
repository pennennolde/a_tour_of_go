package main

import "fmt"

var i, j int = 1, 2
var k = 3
// 関数の外ではvar, funcなどのキーワードではじまる宣言が必要で := は使えない

func main() {
	// var c, python, java = true, false, "no!"
	c, python, java := true, false, "no!"
	// 関数の中ではvar宣言の代わりに := を利用できる

	fmt.Println(i, j, k, c, python, java)
}

