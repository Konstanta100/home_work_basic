package main

import (
	"fmt"
	"sync"
)

func main() {
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
	fmt.Println("Результат счётчика:", sum)
}

func handle(sum *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	*sum++
	fmt.Println("Итерация счётчика:", *sum)
}
