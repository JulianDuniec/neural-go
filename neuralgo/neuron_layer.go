package neuralgo

type NeuronHiddenLayer struct {
	neurons []*Neuron
}

func (this *NeuronHiddenLayer) NumberOfNeurons() int {
	return len(this.neurons)
}

func (this *NeuronHiddenLayer) ProcessInput(input []float64) []float64 {
	numberOfNeurons := len(this.neurons)
	result := make([]float64, numberOfNeurons)

	for i := 0; i < numberOfNeurons; i++ {
		result[i] = this.neurons[i].ProcessInput(input)
	}

	return result
}
