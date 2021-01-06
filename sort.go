package main

import (
	"fmt"
	"sort"
)

// TagWeights 方便排序
type TagWeights struct {
	Tag    string
	Weight float32
}

// TagWeightsList 自定义类型
type TagWeightsList []TagWeights

func (m TagWeightsList) Len() int           { return len(m) }
func (m TagWeightsList) Less(i, j int) bool { return m[i].Weight > m[j].Weight }
func (m TagWeightsList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func sortMain() {
	var mergeList TagWeightsList
	mergeList = append(mergeList, TagWeights{"1", 1})
	mergeList = append(mergeList, TagWeights{"2", 5})
	mergeList = append(mergeList, TagWeights{"2", 3})
	mergeList = append(mergeList, TagWeights{"2", 0})
	mergeList = append(mergeList, TagWeights{"2", 1})

	sort.Sort(mergeList)

	for _, value := range mergeList {
		fmt.Println(value)
	}
}
