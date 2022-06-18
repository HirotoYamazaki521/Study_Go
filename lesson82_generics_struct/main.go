package main

import "fmt"

//generics_struct
//構造体

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

//メソッドは新たに型パラメータを持つことができない
/*
func (t T[A,B,C]) m[U any](x U){
}
*/
func (t T[A, B, C]) m(x int) {
}

func main() {
	var t T[int, []*int, *int]
	fmt.Printf("A: %T, B: %T, C:%T\n", t.a, t.b, t.c)
}
