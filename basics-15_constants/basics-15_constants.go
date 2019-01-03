package main

import "fmt"

// 定数（constant）は、constキーワードを使って宣言する（変数のように省略は不可）
// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使える
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Hello", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

