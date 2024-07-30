package main

import (
	// basictask1 "IrmawanAriel/GolangBasic/BasicTask1"

	basictask2 "IrmawanAriel/GolangBasic/BasicTask2"
	"sync"
	// "fmt"
)

var wg sync.WaitGroup

func main() {
	//nomor 1
	// fmt.Printf("%.2f", basictask1.Pembulatan(4.37))

	// //nomor 2
	// var user basictask1.Bilangan
	// user.N = 9
	// fmt.Println("prima: ", user.Prima())
	// fmt.Println("ganjil: ", user.Ganjil())
	// fmt.Println("genap: ", user.Genap())
	// fmt.Println("Fibonacci: ", user.Fibonacci())

	// // nomor 3
	// persegi := basictask1.Kubus{Sisi: 5}
	// result := []basictask1.Hitung{persegi}
	// for _, result := range result {
	// 	fmt.Printf("bentuk: %T\n", result)
	// 	fmt.Printf("Luas: %.2f\n", result.Luas())
	// 	fmt.Printf("Keliling: %.2f\n", result.Keliling())
	// 	fmt.Printf("Volume: %.2f\n", result.Volume())
	// }

	// // laathian
	// wg.Add(1)
	// x := basictask1.SetTotal(2)
	// y := basictask1.BiayaPengiriman(8)
	// go basictask1.Tampilkan(x, y)
	// wg.Wait()

	// nomor 1 task 2
	// a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	// chn := make(chan int)
	// // go basictask2.Sum(a[:len(a)/2], chn)
	// go basictask2.Sum(a[len(a)/2:], chn)
	// total := <-chn
	// fmt.Println("Total:", total)

	//nomor2 task 2
	ch := make(chan int, 10)
	basictask2.Wg.Add(2)
	go basictask2.Fibonacci(ch)
	go basictask2.GanjilGenap(ch)
	basictask2.Wg.Wait()

	//nomor 3 task 2

}
