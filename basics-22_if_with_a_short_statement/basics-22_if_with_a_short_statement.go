package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// Golangのif文では、for文のように条件の前に簡単なステートメントを書くことができる
	// ここで宣言された変数はifのスコープ内だけで有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // 末尾に余分に , を入れても大丈夫
	)
}

