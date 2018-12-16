package main

import (
	"fmt"
	"math"
)

func main() {

//	fmt.Println(math.pi)

/* 
 * 上記のまま go build すると以下エラーが出る
 *
 * # a_tour_of_go/basics-3_exported_names
 * ./basics-3_exported_names.go:9:14: cannot refer to unexported name math.pi
 * ./basics-3_exported_names.go:9:14: undefined: math.pi
 * 
 * パッケージで定義されている大文字からはじまる名前はインポート先から参照できる。
 */

	fmt.Println(math.Pi)
}

