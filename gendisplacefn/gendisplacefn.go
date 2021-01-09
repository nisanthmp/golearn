package main

import "fmt"

func main() {
	var a, v0, s0, t float64

	fmt.Print("Enter acceleration (in meters per second squared): ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity (in meters per second): ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement (in meters): ")
	fmt.Scan(&s0)

	fmt.Print("Enter time (in seconds): ")
	fmt.Scan(&t)
	fmt.Print("\n")

	fn := genDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %f seconds is %f meters\n", t, fn(t))
}

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		s := 0.5*a*t*t + v0*t + s0
		return s
	}
	return fn
}
