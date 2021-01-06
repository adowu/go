package main

import (
	"fmt"
	"reflect"
	"strings"
)

func simpleArrayMain() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func sliceMain() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:]
	fmt.Println(s)
}

func sliceComplicateMain() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[:2]
	b := names[1:3]
	fmt.Println(a, b)
	c := b
	b[0] = "Adowu"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println(c)
}

func complicateMain() {
	q := []int{2, 3, 5, 7, 11}
	fmt.Println(q)

	r := []bool{true, false, true, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
		{11, false},
	}
	fmt.Println(s)
}

func indexSliceMain() {
	s := []int{3, 5, 7, 11, 13}
	fmt.Println(s)
	fmt.Printf("s type:%T\n", s)
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

	a := [2]int{1, 2}
	fmt.Printf("%T\n", a)
	fmt.Println(reflect.TypeOf(a))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func dynamicSliceMain() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// 切片的容量是从它的第一个元素开始数，到其 底层数组 元素末尾的个数。
	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func defaultSliceMain() {
	var s []string
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func interestSliceMain() {
	a := make([]int, 5)
	// len=5 cap=5 [0 0 0 0 0]
	printSlice(a)

	b := make([]int, 0, 5)
	// len=0 cap=5 []
	printSlice(b)

	c := b[:2]
	// len=2 cap=5 [0 0]
	printSlice(c)

	d := c[2:5]
	// len=3 cap=3 [0 0 0]
	// 为什么这个cap 为 3 主要是left 的index 从一开始的 0 到了2 的位置，到末尾只有 3 个位置了
	// 简单点记就是根据左index 来计算
	printSlice(d)
}

func kindsSliceMain() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func confusedSliceMain() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	// len=5 cap=6 [0 1 2 3 4]
	// s = append(s, 2, 3, 4, 5)
	// len=6 cap=6 [0 1 2 3 4 5]
	// 可以看出来多给了一个位置 s = append(s, 2, 3, 4)
	printSlice(s)

}
