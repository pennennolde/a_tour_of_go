package main

import "fmt"

func main() {
	fmt.Println("counting...")

	for i:=0; i<10; i++ {
		// deferに複数の関数を渡すとstackされる。その実行はLIFO（last-in-first-out）の順になる
		defer fmt.Println(i)
	}

	fmt.Println("done.")
}

