package basictask2

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func Fibonacci(ch chan<- int) {
	defer Wg.Done()

	a, b := 0, 1
	for i := 0; i < 10; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func GanjilGenap(ch <-chan int) {
	defer Wg.Done()

	for num := range ch {
		if num%2 == 0 {
			fmt.Printf("Genap: %d\n", num)
		} else {
			fmt.Printf("Ganjil: %d\n", num)
		}
	}
}
