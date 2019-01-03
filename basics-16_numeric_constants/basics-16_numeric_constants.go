package main

import "fmt"

// 数値の定数は、高精度な値にできる
// 型のない定数は（型を指定して宣言することもできる）、その状況によって必要な型を取ることになる
// intは64-bitの整数を保持できるが、それでは足りない場合に定数を活用する
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) はintの上限を超えてしまうのでエラーになる
}

