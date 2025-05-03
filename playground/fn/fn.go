package fn

import (
	"fmt"
	"math"
)

func SayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func SayBye(n string) {
	fmt.Printf("Good bye %v \n", n)
}

func CycleNames(n []string, f func(string)) {
	for _, val := range n {
		f(val)
	}
}

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}