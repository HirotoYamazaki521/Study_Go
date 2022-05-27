package main

import "fmt"

//型
//浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	//暗黙的な定義で書いた場合は自動でfloat64になる
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	//float32は基本的にはあんまり使わない
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	/*
		０で割ると
		+Inf 正の無限大
		-Inf 負の無限大
		NaN 非数　特殊な値を保持
	*/
}
