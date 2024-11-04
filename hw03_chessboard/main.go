package main

import (
	"fmt"
)

func main() {
	var width, height int

	fmt.Printf("Введите ширину доски: ")
	fmt.Scanln(&width)
	fmt.Printf("Введите высоту доски: ")
	fmt.Scanln(&height)

	var result, first, last string

	for i := 0; i < height; i++ {
		if first == "#" {
			first = " "
		} else {
			first = "#"
		}
		last = first

		for j := 0; j < width; j++ {
			if last == "#" {
				last = " "
			} else {
				last = "#"
			}

			result += last
		}

		result += "\n"
	}

	fmt.Printf("%s", result)
}
