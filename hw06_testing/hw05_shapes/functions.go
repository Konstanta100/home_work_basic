package shapes

import (
	"errors"
	"fmt"
)

func calculateArea(s any) (string, string, error) {
	shape, err := s.(Shape)

	if !err {
		return "", "", errors.New("переданный объект не является фигурой")
	}

	return fmt.Sprintf("Площадь: %.2f\n\n", shape.Area()), shape.Description(), nil
}
