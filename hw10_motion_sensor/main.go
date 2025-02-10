package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Start:")
	chSensorBuffer := make(chan int)
	chResult := make(chan float64)

	go generateSensorBuffer(chSensorBuffer)
	go generateResult(chResult, chSensorBuffer)

	for result := range chResult {
		fmt.Printf("Result process: %.2f\n", result)
	}

	fmt.Println("End program")
}

func generateSensorBuffer(chBuffer chan<- int) {
	chTimer := time.After(56 * time.Second)

	for {
		select {
		case <-chTimer:
			close(chBuffer)
			fmt.Println("Timer is out")
			return
		case chBuffer <- rand.Intn(100): //nolint:gosec
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func generateResult(chResult chan<- float64, chBuffer <-chan int) {
	defer close(chResult)
	var count int
	var sum float64
	for buffer := range chBuffer {
		count++
		sum += float64(buffer)

		if count == 10 {
			average := calculateAverage(sum, count)
			chResult <- average
			count = 0
			sum = 0
		}
	}
}

func calculateAverage(sum float64, count int) float64 {
	return sum / float64(count)
}
