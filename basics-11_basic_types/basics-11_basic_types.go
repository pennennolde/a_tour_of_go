package main

import (
	"fmt"
	"math/cmplx"
)

/*
 * Golangの基本型（組み込み型）
 *
 * bool
 *
 * string
 *
 * // int, uint, uintptr型は、32-bitシステムでは32bit、64-bitシステムでは64bit
 * // uint = unsigned int (符号なし整数 - 0以上の整数値)
 * int  int8  int16  int32  int64
 * uint uint8 uint16 uint32 uint64 uintptr
 *
 * byte // uint8 の別名
 *
 * rune // int32 の別名
 *      // Unicode のコードポイントを表す（文字そのものを表す）
 *
 * float32 float64
 *
 * complex64 complex128
 *
 */

// 変数宣言はfactored宣言が可能
var (
	ToBe    bool      = false
	MaxInt uint64     = 1<<64 - 1  // <<は、右辺の値分だけ左辺を算術左シフトする演算子
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

