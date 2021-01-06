package main

import (
	"fmt"
	"math"
)

// Cnice X and Y
type Cnice struct {
	X, Y float64
}

// Abs sqrt
func (v Cnice) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// normalAbs 下面这个就是正常清空下我们的书写方式
func normalAbs(v Cnice) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale pointer function
func (v *Cnice) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func normalScale(v Cnice, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func simple1Main() {
	v := Cnice{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(normalAbs(v))

}

func xiMain() {
	v := Cnice{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}
