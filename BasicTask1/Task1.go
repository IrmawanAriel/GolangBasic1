package basictask1

import (
	"math"
)

func Pembulatan(n float64) float64 {
	return math.Round(n*10) / 10
}
