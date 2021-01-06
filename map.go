package main

import (
	"fmt"
	"strings"
)

type Nice struct {
	Lat, Long float64
}

func simpleMain() {
	m = make(map[string]Nice)
	m["Bell Labs"] = Nice{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

}

var m = map[string]Nice{
	"Bell Labs": {
		40.123, -43.223,
	},
	"Google": {
		23.234, -12.124,
	},
}

func moreMain() {
	fmt.Println(m)
}

func operationMain() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println(m["Answer"])

	m["Answer"] = 48
	fmt.Println(m["Answer"])

	delete(m, "Answer")
	fmt.Println(m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("Value:", v, "Present?", ok)

}

// WordCount count
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	splits := strings.Split(s, " ")
	for _, val := range splits {
		if _, ok := result[val]; ok {
			result[val]++
		} else {
			result[val] = 1
		}
	}
	return result
}

func oomain() {
	s := "i like you! but you dont like me."
	fmt.Println(WordCount(s))
	result := WordCount(s)
	for k, v := range result {
		fmt.Println(k)
		fmt.Println(v)
	}

	a := []string{"1"}
	fmt.Println(a)
	b := make(map[string][]string)
	b["a"] = append(b["a"], "c")
	fmt.Println(b)

	items := make(map[string]interface{})
	items["a"] = 1
	if v, ok := items["a"].(int); ok && v == 0 {
		fmt.Println("ok,oooops")
	} else {
		fmt.Println("Noooooooo")
	}

}
