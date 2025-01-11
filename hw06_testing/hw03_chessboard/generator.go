package chessboard

import (
	"errors"
)

func generateBoard(width, height int) (string, error) {
	var result, last string

	if width <= 0 || height <= 0 {
		return "", errors.New("ширина и высота должны быть больше 0")
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (j+i)%2 == 0 {
				last = " "
			} else {
				last = "#"
			}

			result += last
		}

		result += "\n"
	}

	return result, nil
}
