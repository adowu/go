package main

import (
	"fmt"
	"reflect"
	"sort"
)

// BodyWrapper .
type BodyWrapper struct {
	Bodys []interface{}
	by    func(p, q *interface{}) bool
}

// SortBodyBy .
type SortBodyBy func(p, q *interface{}) bool

// Len .
func (wrapper BodyWrapper) Len() int { return len(wrapper.Bodys) }

// Less .
func (wrapper BodyWrapper) Less(i, j int) bool {
	return wrapper.by(&wrapper.Bodys[i], &wrapper.Bodys[j])
}

// Swap .
func (wrapper BodyWrapper) Swap(i, j int) {
	wrapper.Bodys[i], wrapper.Bodys[j] = wrapper.Bodys[j], wrapper.Bodys[i]
}

// SortBody .
func SortBody(bodys []interface{}, by SortBodyBy) {
	sort.Sort(BodyWrapper{bodys, by})
}

// Person .
type Person struct {
	Name       string
	CreateTime string
}

func sssmain() {
	results := []interface{}{}
	u1 := Person{
		Name:       "A",
		CreateTime: "2018-04-01",
	}
	u2 := Person{
		Name:       "B",
		CreateTime: "2018-03-01",
	}
	results = append(results, u1)
	results = append(results, u2)
	fmt.Println(results)
	SortBody(results, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName("CreateTime")
		v = reflect.ValueOf(*q)
		j := v.FieldByName("CreateTime")
		return i.String() < j.String()
	})
	fmt.Println(results)
}
