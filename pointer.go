package main

import "fmt"

func pointerMain() {
	i, j := 42, 2701
	p := &i
	// 只输出p的话，这时候就是一个地址了
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(*p)

}
