package main

import (
	basictask2 "IrmawanAriel/GolangBasic/BasicTask2"
)

// var wg sync.WaitGroup

func main() {
	// // nomor 1
	// fmt.Printf("%.2f", basictask1.Pembulatan(4.37))

	// // nomor 2
	// var user basictask1.Bilangan
	// user.N = 9
	// fmt.Println("prima: ", user.Prima())
	// fmt.Println("ganjil: ", user.Ganjil())
	// fmt.Println("genap: ", user.Genap())
	// fmt.Println("Fibonacci: ", user.Fibonacci())

	// // nomor 3
	// var persegi basictask1.Hitung       // insialisasi var kedalam interface hitungnya
	// persegi = basictask1.Kubus{Sisi: 5} // jika var tersebut di inisialisasi kedalam struct yang tidak sesuai denganinterface maka error
	// fmt.Println("ini luas permukaan :", persegi.Luas())
	// fmt.Println("ini keliling:", persegi.Keliling())
	// fmt.Println("ini volume :", persegi.Volume())

	// // latihan
	// wg.Add(1)
	// x := basictask1.SetTotal(2)
	// y := basictask1.BiayaPengiriman(8)
	// go basictask1.Tampilkan(x, y)
	// wg.Wait()

	// nomor 1 task 2
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)
	basictask2.Wg.Add(1)
	go basictask2.Sum(a[len(a)/2:], chn)
	basictask2.Wg.Wait()
	close(chn)

	// // nomor 2 task 2
	// ch := make(chan int, 10)
	// basictask2.Wg.Add(2)
	// go basictask2.Fibonacci(ch)
	// go basictask2.GanjilGenap(ch)
	// basictask2.Wg.Wait()

	// //nomor 3 task 2
	// ch := make(chan int, 10) // banner
	// basictask2.Wg.Add(2)
	// go basictask2.Fibonacci(ch) // routine
	// go basictask2.GanjilGenap(ch) // routine
	// basictask2.Wg.Wait()

}
