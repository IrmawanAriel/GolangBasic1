package basictask1

import "math"

type Bilangan struct {
	N int
}

func (angka Bilangan) Prima() []int {
	var result []int

	for i := 2; i <= angka.N; i++ {
		valid := true
		sqrtI := int(math.Sqrt(float64(i)))

		for j := 2; j <= sqrtI; j++ {
			if i%j == 0 {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, i)
		}
	}

	return result
}

func (angka Bilangan) Ganjil() []int {
	var resultGanjil []int

	for i := 0; i <= angka.N; i++ {
		if i%2 != 0 {
			resultGanjil = append(resultGanjil, i)
		}
	}
	return resultGanjil
}

func (angka *Bilangan) Genap() []int {
	var resultGenap []int

	for i := 0; i <= angka.N; i++ {
		if i%2 == 0 {
			resultGenap = append(resultGenap, i)
		}
	}
	return resultGenap
}

func (angka Bilangan) Fibonacci() []int {

	resultFib := make([]int, angka.N+1)
	resultFib[0] = 0
	resultFib[1] = 1

	for i := 2; i <= angka.N; i++ {
		resultFib[i] = resultFib[i-1] + resultFib[i-2] // looping kebelakang
	}

	return resultFib

}
