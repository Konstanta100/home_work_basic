package main

func main() {
	result := binarySearch([]int{10, 20, 30, 35, 50, 56, 75, 90, 100, 110, 200, 201, 305, 708, 805}, 805)
	println(result)
}

func binarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2
		element := array[mid]

		if element == target {
			return mid
		}

		if element > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
