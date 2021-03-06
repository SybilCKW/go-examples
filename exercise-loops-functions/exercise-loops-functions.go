// See: https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i:=0 ; i<5; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}
