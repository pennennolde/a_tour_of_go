package main

import "fmt"

func main() {
	i, j := 42, 2701

	// &i → 値iが格納されているメモリアドレスを指すポインタ
	// *p → ポインタpが指しているメモリアドレスに格納されている値

	var p *int	// ポインタの定義 (ポインタは *オペレータによってそれが指しているメモリアドレスに格納されている値を参照できるint型の値)

	p = &i			// point to i （iのポインタを引き出してポインタ値をpに格納）
	fmt.Println("*p: ", *p)	// read i through the pointer （ポインタが指しているメモリアドレスに格納されている値を引き出す）
	fmt.Println("p: ", p)
	k := &p			// ポインタのポインタ？？
	fmt.Println("k: ", k)
	*p = 21			// set i through the pointer （ポイントからi参照して、値を格納する）
	fmt.Println("i: ", i)

	p = &j
	*p = *p / 37		// divide j through the pointer
	fmt.Println("j: ", j)
}

