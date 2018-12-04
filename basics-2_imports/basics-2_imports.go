package main

// パッケージのインポートは括弧でグループ化して記述する
//  → 推奨
import (
	"fmt"
	"math"
)
// 並べて書くことも可
// import 'fmt"
// import "math"

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

