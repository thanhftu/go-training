package main

import "fmt"

type T1 struct {
	field1 string
}

func (t *T1) M() {
	t.field1 = t.field1 + t.field1
}

type T2 struct {
	field2 string
	T1
}

type I interface {
	M()
}

func main() {
	var i I = &T2{"foo", T1{"bar"}}
	i.M()
	fmt.Println(i.(*T2).field1)
}
