package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}	// has type Vertex
	p  = &Vertex{1, 2}	// has type *Vertex (*int ではなく *Vertex）
	v2 = Vertex{X: 1}	// Name: 構文を使ってフィールドの一部だけを初期化できる。指定されなかった他のフィールドは暗に初期化される（今回は Y:0 になる）
	v3 = Vertex{}		// 暗に X:0 and Y:0 と初期化される
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Println(p)		// *int みたいなメモリアドレス的な値は出力されないが、
	fmt.Println(*p)		// &オペレータを使うとポインタが指しているstructが返るので、pはちゃんとポインタの役割をしているとわかる
}

