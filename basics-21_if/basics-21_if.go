package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// for文同様、if文も()は不要
	if x < 0 {
		// sqrtを再帰的に使用
		return sqrt(-x) + "i" // sqrt(-x) は結果的にstringで返ってくるので "string + string" になる
	}
	return fmt.Sprint(math.Sqrt(x)) // 結果を文字列で返す
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

