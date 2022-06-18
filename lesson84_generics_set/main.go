package main

import "fmt"

//generics_set
//keyとvalueをもつmap型のような構造をしている

//comparable　要素の比較が可能であるという型制約をかけれる
type Set[T comparable] map[T]struct{} //mapのkeyには比較可能な型を設定する必要がある

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, v := range xs {
		//s[v] = struct{}{}
		s.Add(v)
	}

	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {
	s := NewSet(1, 2, 3)
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)

}
