package main

import "fmt"

var x = 0

func example3() {
	x = 10
}

func example2() int {
	defer example3()
	return x
}

func example() {
	if x+=example2()+1; x==11 {
		fmt.Println("example2()の返り値（Golangでは返り値は基本的に値。参照ではなく）が決定してコピーが作成され、deferに渡した関数が実行され、その後返り値を使った計算や代入の処理がされていく")
		fmt.Println("（example2()のreturn値0が決定してコピーが作成され、deferに渡したexample3()が実行され、x+=0+1 が計算される）")
	}
}

func main() {
	// deferへ渡した関数の実行を、呼び出し元の関数がreturnするまで遅延させる
	// deferに渡した関数の引数はすぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで実行されない
	// deferに渡すのは関数じゃないとダメ、計算式とかはダメ
	defer fmt.Println("完了")
	fmt.Println("Hello")

	example()
}

