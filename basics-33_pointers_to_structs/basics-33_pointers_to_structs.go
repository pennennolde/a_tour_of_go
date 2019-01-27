package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// フィールドXを持つstructのポインタpがあるとき、フィールドXにアクセスするには (*p).X のように書くことができるが、
	// Golangでは簡単に p.X と書くこともできる（structを指すポインタからフィールドを参照できる）
	p.X = 1e9
	fmt.Println(v)
}

