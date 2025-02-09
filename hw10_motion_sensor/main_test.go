package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSensorBuffer(t *testing.T) {
	chBuffer := make(chan int)
	done := make(chan struct{})

	go func() {
		generateSensorBuffer(chBuffer)
		close(done)
	}()

	timeout := time.After(61 * time.Second)
	var values []int

	for {
		select {
		case v, ok := <-chBuffer:
			if !ok {
				if len(values) == 0 {
					t.Error("Expected some values, got none")
				}
				return
			}
			values = append(values, v)
		case <-timeout:
			t.Error("Test timed out")
			return
		case <-done:
			return
		}
	}
}

func TestGenerateResult(t *testing.T) {
	chSensorBuffer := make(chan int)
	chResult := make(chan float64)

	go generateResult(chResult, chSensorBuffer)

	go func() {
		for i := 0; i < 10; i++ {
			chSensorBuffer <- 10
		}
		close(chSensorBuffer)
	}()

	select {
	case result := <-chResult:
		if result != 10.0 {
			t.Errorf("Expected average to be 10.0, got %v", result)
		}
	case <-time.After(1 * time.Second):
		t.Error("Test timed out")
	}
}

func TestCalculateAverageMap(t *testing.T) {
	testCases := []struct {
		name     string
		sum      float64
		count    int
		expected float64
	}{
		{"avg_1", 100.0, 10, 10.0},
		{"avg_2", 50.0, 5, 10.0},
		{"empty", 0.0, 10, 0.0},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := calculateAverage(tC.sum, tC.count)
			assert.Equal(t, tC.expected, result)
		})
	}
}
