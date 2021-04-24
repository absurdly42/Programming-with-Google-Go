package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Введите ускорение:")
	fmt.Scanln(&a)
	var v0 float64
	fmt.Print("Введите начальную скорость:")
	fmt.Scanln(&v0)
	var s0 float64
	fmt.Print("Введите начальное смещение:")
	fmt.Scanln(&s0)
	var t float64
	fmt.Print("Введите время:")
	fmt.Scanln(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1/2*a*t*t + v0*t + s0
	}
}
