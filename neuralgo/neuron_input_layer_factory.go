package neuralgo

/*
   Factory
*/
func NewNeuronInputLayer(numberOfInputs int, activationFunction ActivationFunction) *NeuronInputLayer {
	return &NeuronInputLayer{
		NewNeuronLayer(numberOfInputs, 1, activationFunction),
	}
}
