package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forOrigin() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt((x)))
}

func forSimple() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func ifOrigin() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func ifSimple() string {
	return fmt.Sprint(
		pow(3, 3, 20),
	)

}

func ifSimpleMain() {
	result := ifSimple()
	fmt.Println(result)
}

func caseMain() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}

func switchWithNoMain() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func deferMain() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func deferLoopMain() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
