package main

import "fmt"

//interface
//最もポピュラーな使い方、異なる型に共通の性質を付与する

type Stringfy interface {
	//ToStringという文字列を返すメソッドをStringfyという型でまとめる
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

//出力系のメソッドを用意
func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age =%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model =%v", c.Number, c.Model)
}

func main() {
	//通常はPersonとCarは異なるデータ型なのでスライスでまとめることはできないが、
	//共通の性質を持つStringfyという型でまとめて管理できる
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	//goの厳密な型のシステムに柔軟性を与えることができる
}
