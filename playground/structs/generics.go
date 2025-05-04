package structs

import (
	"math"
)

type GShape interface {
	GArea() float64
}

type GCircle struct { Radius float64 }
type GRectangle struct { Width, Height float64 }

func (c GCircle) GArea() float64 { return math.Pi * c.Radius * c.Radius }

func (r GRectangle) GArea() float64 { return r.Width * r.Height }

func SumAreas[T GShape](shapes []T) float64 {
	var total float64
	for _, s := range shapes {
		total += s.GArea()
	}
	return total
}