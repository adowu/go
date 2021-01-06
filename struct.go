package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func simpleStructMain() {
	fmt.Println(Vertex{1, 2})
}

func pointerStructMain() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 4}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func varStructMain() {
	fmt.Println(v1, v2, v3, *p)
}
