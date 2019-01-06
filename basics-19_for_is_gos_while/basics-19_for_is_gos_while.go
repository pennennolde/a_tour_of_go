package main

import "fmt"

func main() {
	sum := 1
	// ; を省略して書くこともできる → 通例のwhile文になる
	//  Golangにwhile文はなく、for文でこのように書く
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

