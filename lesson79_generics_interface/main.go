package main

import (
	"fmt"
	"strconv"
)

//Generics interface
//型パラメータに型の制約をもたせる

type Stringer interface {
	String() string
}

//func f[T fmt.Stringer](xs []T) []string {
func f[T Stringer](xs []T) []string { //interfaceを満たす型パラメータのみを指定する制約
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String()) //fmt.Stringerを満たす型制約をかけているので、この中ではx.String()を使うことができる。
	}
	return result
}

type MyInt int

//MyInt型はString()メソッドを持っていてfmt.Stringerを満たしているので、
//fの関数の型パラメータとして指定することができるようになる。
func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))
}
