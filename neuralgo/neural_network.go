package neuralgo

type NeuralNetwork struct {
	layers []NeuronLayer
}

func (this *NeuralNetwork) ProcessInput(input []float64) []float64 {
	for i := 0; i < len(this.layers); i++ {
		input = this.layers[i].ProcessInput(input)
	}
	return input
}
