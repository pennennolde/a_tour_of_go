package main

import (
	"fmt"
	"runtime"
)

/*
 * Golangのswitch文は、あてはまったcaseのみ実行し、それにつづくcaseは実行されない
 *  → breakを書く必要はない、自動的に実行されるということ
 * また、caseは定数である必要はなく、関係する値は整数である必要もない
 *  → if文のような条件式を書ける
 */
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	case "nacl":
		fmt.Println("Nacl.")
	default:
		fmt.Printf("%s.", os)
	}
}

