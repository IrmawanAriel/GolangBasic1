package basictask1

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func SetTotal(n int) int {

	var total int = 9
	subtotal := total + n

	return subtotal

}

func BiayaPengiriman(s int) int {

	var hargajarak = s * 1000
	return hargajarak
}

func Tampilkan(n int, s int) {
	defer wg.Done()

	fmt.Print("total harga : ", n, "\ntotal pengiriman :", s, "\ntotal keseluruhan : ", s+n)
}
