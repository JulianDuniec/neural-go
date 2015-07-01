package neuralgo

import (
	"math"
)

func Sigmoid(input float64) float64 {
	return 1 / (1 + math.Pow(math.E, -input))
}
