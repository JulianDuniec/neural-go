package neuralgo

import (
	"fmt"
)

type Neuron struct {
	weights            []float64
	activationFunction ActivationFunction
}

func (this *Neuron) ProcessInput(input []float64) float64 {

	inputLen, weightsLen := len(input), len(this.weights)-1

	// Guard agains sending too many/few inputs
	if inputLen != weightsLen {
		panic(fmt.Sprintf("Neuron recieved input out of bounds, can handle %i, got %i", weightsLen, inputLen))
	}

	var activation float64

	activation = 0

	for i := 0; i < len(input); i++ {
		activation += input[i] * this.weights[i]
	}

	activation += this.weights[weightsLen]

	return this.activationFunction(activation)
}
