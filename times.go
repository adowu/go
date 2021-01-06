package main

import (
	"fmt"
	"time"
)

const gap int = 10 * 24 * 60 * 60
const UINT_MAX uint = ^uint(0)
const INT_MAX int = int(^uint(0) >> 1)

func ttttmain() {
	var filterCount int
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(int(now))

	fmt.Println(gap)
	fmt.Println(INT_MAX)
	fmt.Println(UINT_MAX)
	fmt.Println(filterCount)
	utc := time.Now().UTC()
	time.Sleep(10000)
	time.Now().UTC().Sub(utc).Microseconds()
	啊 := time.Now().Local().Format("2006-01-02")
	fmt.Println(啊)
	expire := time.Duration(3600 * time.Second).Seconds()
	fmt.Println(expire)
	fmt.Println(time.Second)
	fmt.Println(time.Duration(400) * time.Millisecond)

}
