package main

import (
	"fmt"
	"math"
)

// Goでの型変換は明示的は変換が必要
func main() {
	var x, y float64 = 3, 4
	f := math.Sqrt(x*x + y*y)
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	var i int = 42
	var f2 float64 = float64(i)
	var u uint = uint(f2)
	fmt.Println(i, f2, u)

	// より簡潔に記述できる
	ii := 42
	ff := float64(ii)
	uu := uint(ff)
	fmt.Println(ii, ff, uu)
}

