package main

import (
	"fmt"
	"sort"
)

//sort

//スライス等をソートする機能
/*
func main() {
	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s)
}
*/

//SliceとSclisetable
/*
type Entry struct {
	Name  string
	Value int
}

func main() {
	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	//Slice
	//第1引数にスライス、第2引数に条件式
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name }) //名前を昇順にそーと

	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value }) //名前順にした後に数値順にするが、数値順にしかなっていない

	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")

	//SliceStable
	//Sliceとの違いは安定ソートであるか
	//ソート途中の各状態に常に従位の位置関係を保ってること言う
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")
}
*/

//CustomSort
type Entry struct {
	Name  string
	Value int
}
type List []Entry

/*
type Interface interface{
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i,j int) bool
	// Swap swaps the elements with indexes i and j.const
	Swap(i,j int)
}
*/

func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//ここをカスタムする
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name) //Name順にする
	} else {
		return (l[i].Value < l[j].Value) //Value順にする
	}
}

func main() {

	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e) //ListにEntry型の値を追加していくもの
	}

	//Sort
	sort.Sort(lt)

	fmt.Println(lt)

	//Reverse	//ソートした内容をそのままひっくり返してソート
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)

}
