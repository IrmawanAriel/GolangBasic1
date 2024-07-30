package main

import (
	basictask1 "IrmawanAriel/GolangBasic/BasicTask1"
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// fmt.Printf("%.2f", basictask1.Pembulatan(4.37))
	var user basictask1.Bilangan
	user.N = 9
	fmt.Println("prima: ", user.Prima())
	fmt.Println("ganjil: ", user.Ganjil())
	fmt.Println("genap: ", user.Genap())
	fmt.Println("Fibonacci: ", user.Fibonacci())

	// wg.Add(1)

	// x := basictask1.SetTotal(2)
	// y := basictask1.BiayaPengiriman(8)

	// go basictask1.Tampilkan(x, y)
	// wg.Wait()

}
