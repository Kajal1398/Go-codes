package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*(math.Pow(t, 2)*0.5) + (v0 * t) + (s0))
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Println("enter values for acceleration, initial velocity, and initial displacement and time")
	fmt.Scan(&a, &v0, &s0, &t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement: %f\n", fn(t))

}
