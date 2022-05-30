package main

import (
	"fmt"
	"strconv"
)

//if
//条件分岐
//エラーハンドリング

func main() {
	var s string = "A"

	//i, _ := strconv.Atoi(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", i)
}
