package entity

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Description() string {
	return fmt.Sprintf("Круг: радиус %.2f", c.Radius)
}
