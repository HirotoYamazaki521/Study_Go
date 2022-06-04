package main

import "fmt"

//interface
//fmt.Stringer
//文字列を返すメソッドのStringのみが実装されているインターフェース

/*
type Stringer interface{
	String() string
}
*/

type Point struct {
	A int
	B string
}

//表示形式のカスタマイズ
func (p *Point) String() string {
	return fmt.Sprintf("<<%v,%v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
}
