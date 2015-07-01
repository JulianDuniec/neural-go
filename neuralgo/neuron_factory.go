package neuralgo

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
   Factory
*/
func NewNeuron(numberOfInputs int, activationFunction ActivationFunction) *Neuron {
	return &Neuron{
		weights:            buildWeights(numberOfInputs),
		activationFunction: activationFunction,
	}
}

func buildWeights(numberOfInputs int) []float64 {

	weights := make([]float64, numberOfInputs+1)

	for i := 0; i < len(weights); i++ {
		weights[i] = (random.Float64() * 2) - 1
	}

	return weights
}
