package entity

import "fmt"

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Description() string {
	return fmt.Sprintf("Прямоугольник: ширина %.2f, высота %.2f", r.Width, r.Height)
}
