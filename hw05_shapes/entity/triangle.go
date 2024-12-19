package entity

import "fmt"

type Triangle struct {
	Height float64
	Width  float64
}

func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}

func (t Triangle) Description() string {
	return fmt.Sprintf("Треугольник: основание %.2f, высота %.2f", t.Width, t.Height)
}
