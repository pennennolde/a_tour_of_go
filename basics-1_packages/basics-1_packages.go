// プログラムの実行はmainパッケージの実行からはじまる
package main

// fmtパッケージとmath/randパッケージをインポート
import (
	"fmt"
	"math/rand"
)

// パッケージ名はインポートパスの最後の要素と同じ名前になる
// math/rand -> rand
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

