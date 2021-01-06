package main

import "fmt"

func emain() {
	a := map[string]interface{}{
		"name": "wushaojun",
		"id":   2,
	}

	if value, ok := a["age"].(int); ok {
		fmt.Println(value)
	} else {
		fmt.Println(ok)
	}
}
