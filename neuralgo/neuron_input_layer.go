package neuralgo

type NeuronInputLayer struct {
	*NeuronHiddenLayer
}

func (this *NeuronInputLayer) ProcessInput(input []float64) []float64 {
	numberOfNeurons := len(this.neurons)
	result := make([]float64, numberOfNeurons)

	for i := 0; i < numberOfNeurons; i++ {
		result[i] = this.neurons[i].ProcessInput([]float64{input[i]})
	}

	return result
}
