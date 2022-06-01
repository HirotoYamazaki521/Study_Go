package main

import "fmt"

//switch
//型スイッチ
//型アサーション 動的に変数の型をチェっクする

func anything(a interface{}) {
	//fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	case float64:
		fmt.Println(v + 1.5)
	}
}

func main() {
	anything("aaa")
	anything(1)
	anything(1.0)

	var x interface{} = 3
	i, isInt := x.(int) //変数.(復元したい型)
	fmt.Println(i, isInt)
	//fmt.Println(x + 2)

	/*
		f := x.(float64)	//float64では復元できないためランタイムエラー
		fmt.Println(f)
	*/

	f, isFloat64 := x.(float64) //変換が成功すると2つ目の変数にtrue
	fmt.Println(f, isFloat64)
	//2つの返り値を使った型アサーションを使えば、プログラムを強制終了させることなく実行できる

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't Know")
	}

	//switchを使った型アサーションの方が簡単なため、こっちが推奨
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	//上だと復元した値を使えないため、値を使いたい場合
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "stirng")
	}
}
