package main

import (
	"fmt"
	"strings"

	"github.com/wxnacy/wgo/arrays"
)

func arrrayMain() {
	a := make([]int, 5)
	fmt.Println(a)
	// panic: runtime error: slice bounds out of range [:10] with capacity 5
	fmt.Println(a[:10])
}

func switchMain() {
	f := "shaojun"
	switch f {
	case "adowu":
		fmt.Println("adowu")
	case "wushaojun":
		fmt.Println("wushaojun")
	default:
		fmt.Println("None")
	}
}

func forMmain() {
	tags := []int{1, 2, 3, 4, 5}
	for index := range tags {
		fmt.Print(index)
	}
}

func listMain() {
	var mergeList [][2]interface{}
	mergeList = append(mergeList, [2]interface{}{1, "1"})
	mergeList = append(mergeList, [2]interface{}{5, "2"})
	mergeList = append(mergeList, [2]interface{}{3, "2"})
	mergeList = append(mergeList, [2]interface{}{0, "2"})
	mergeList = append(mergeList, [2]interface{}{2, "2"})

}

func ssssMain() {
	a := make([]int, 0)
	fmt.Println(a)
	b := "123"
	c := "qwwerasd"
	index := arrays.Contains(a, 4)
	fmt.Println(index)
	tite := fmt.Sprintf("unseen recall prefix: %s, reqID: %s", b, c)
	fmt.Println(tite)
	d := strings.Split(c, "_")
	fmt.Println(d)
	e := map[string]int{
		"1": 1,
		"2": 2,
	}
	fmt.Println(e)
}

func tttmain() {
	a := make(map[string][]int)
	a["a"] = append(a["a"], 1)
	a["a"] = append(a["a"], []int{2, 34, 5}...)

	fmt.Println(a)

}
