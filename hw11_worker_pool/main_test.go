package main

import (
	"sync"
	"testing"
)

func TestHandler(t *testing.T) {
	const countHandlers = 1000
	var (
		mu  sync.Mutex
		sum int
	)

	wg := sync.WaitGroup{}

	for i := 1; i <= countHandlers; i++ {
		wg.Add(1)
		go handle(&sum, &wg, &mu)
	}

	wg.Wait()

	if sum != countHandlers {
		t.Errorf("wrong number of handlers: got %d, want %d", sum, countHandlers)
	}
}
