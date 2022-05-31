package main

import "fmt"

//panic recover
//例外処理
//プログラムの処理を強制終了させてしまうため、あまり使わないほうが良い

func main() {
	defer func() {
		if x := recover(); x != nil { //panic状態でなければxに値が入る
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Strat")
}

//panicになるとdeferの中身しか実行されないため、
//原則deferとセットで使う
