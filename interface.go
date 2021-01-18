package main

import "fmt"

// I define method M
type I interface {
	M()
}

// T define S
type T struct {
	S string
}

// M implements I
func (t T) M() {
	fmt.Println(t.S)
}

func interfaceMain() {
	var i I = T{"hello"}
	i.M()
}
