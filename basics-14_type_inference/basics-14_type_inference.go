package main

import "fmt"

func main() {
	// 明示的に型を指定せずに変数を宣言する場合（ := か var = のいずれか）、変数の型は右辺の変数から型類推される
	// 右辺の変数が型を持っている場合は、左辺の新しい変数は同じ型になる
	v := ""
	fmt.Printf("v is of type %T\n", v)

	var a int
	b := a // b is an int
	fmt.Printf("type: %T\n", b)

	// 右辺が型を指定しない数値である場合は、その数値の精度に基づいて int, float64, complex128 のいずれかの型になる
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("type: %T\n", i)
	fmt.Printf("type: %T\n", f)
	fmt.Printf("type: %T\n", g)
}

