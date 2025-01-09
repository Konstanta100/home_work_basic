package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArea_MainTableTest(t *testing.T) {
	testCases := []struct {
		name        string
		shape       any
		exception   string
		area        string
		description string
	}{
		{
			name:        "TriangleCase",
			shape:       Triangle{Height: 20, Width: 40},
			exception:   "",
			area:        "Площадь: 400.00\n\n",
			description: "Треугольник: основание 40.00, высота 20.00",
		},
		{
			name:        "RectangleCase",
			shape:       Rectangle{Height: 20, Width: 15},
			exception:   "",
			area:        "Площадь: 300.00\n\n",
			description: "Прямоугольник: ширина 15.00, высота 20.00",
		},
		{
			name:        "CircleCase",
			shape:       Circle{Radius: 15},
			exception:   "",
			area:        "Площадь: 706.86\n\n",
			description: "Круг: радиус 15.00",
		},
		{
			name:        "EmptyCase",
			shape:       nil,
			exception:   "переданный объект не является фигурой",
			area:        "",
			description: "",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			area, description, err := calculateArea(tC.shape)

			assert.Equal(t, tC.area, area)
			assert.Equal(t, tC.description, description)

			if err != nil {
				assert.Equal(t, tC.exception, err.Error())
			}
		})
	}
}
