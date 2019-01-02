package main

import "fmt"

func main() {
	// 変数に初期値を与えずに宣言すると、ゼロ値（zero vaalue）が与えられる
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}

