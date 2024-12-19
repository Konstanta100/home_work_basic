package main

import (
	"errors"
	"fmt"

	"github.com/Konstanta100/home_work_basic/hw05_shapes/entity"
)

func main() {
	triangle := entity.Triangle{Height: 20, Width: 40}
	rectangle := entity.Rectangle{Height: 20, Width: 15}
	circle := entity.Circle{Radius: 15}

	shapes := []any{triangle, rectangle, circle, "test Error", 1}

	for _, shape := range shapes {
		area, description, err := calculateArea(shape)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println(description)
		fmt.Printf("Площадь: %.2f\n\n", area)
	}
}

func calculateArea(s any) (float64, string, error) {
	shape, err := s.(entity.Shape)

	if !err {
		return 0, "", errors.New("переданный объект не является фигурой")
	}

	return shape.Area(), shape.Description(), nil
}
