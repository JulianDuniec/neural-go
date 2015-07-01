package neuralgo

/*
   Factory
*/
func NewNeuronLayer(numberOfNeurons, numberOfInputsPerNeuron int, activationFunction ActivationFunction) *NeuronHiddenLayer {
	neurons := make([]*Neuron, numberOfNeurons)

	for i := 0; i < numberOfNeurons; i++ {
		neurons[i] = NewNeuron(numberOfInputsPerNeuron, activationFunction)
	}
	return &NeuronHiddenLayer{
		neurons: neurons,
	}
}
