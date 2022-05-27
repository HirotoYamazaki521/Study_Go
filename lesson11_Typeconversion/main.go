package main

import (
	"fmt"
	"strconv"
)

//型変換

func main() {
	/*
		//数値の返還
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl64 = %T\n", fl64)

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2)

		fl := 10.5
		i3 := int(fl) //小数点以下を切り捨て
		fmt.Printf("i3 = %T\n", i3)
		fmt.Println(i3)
	*/

	//文字列型から数値型への変換
	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	fmt.Println(i)

}
