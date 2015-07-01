package neuralgo

func NewNeuralNetwork(
	numberOfInputs,
	numberOfOutputs,
	numberOfHiddenLayers,
	numberOfNeuronsPerHiddenLayer int,
	activationFunction ActivationFunction) *NeuralNetwork {

	// Input layer + hidden layers + output layer
	numberOfLayers := numberOfHiddenLayers + 2

	layers := make([]NeuronLayer, numberOfLayers)

	for i := 0; i < numberOfLayers; i++ {
		if i == 0 {
			layers[i] = NewNeuronInputLayer(numberOfInputs, activationFunction)
		} else if i == numberOfLayers-1 {
			layers[i] = NewNeuronLayer(numberOfOutputs, numberOfNeuronsPerHiddenLayer, activationFunction)
		} else if i == 1 {
			layers[i] = NewNeuronLayer(numberOfNeuronsPerHiddenLayer, numberOfInputs, activationFunction)
		} else {
			layers[i] = NewNeuronLayer(numberOfNeuronsPerHiddenLayer, numberOfNeuronsPerHiddenLayer, activationFunction)
		}
	}
	return &NeuralNetwork{
		layers: layers,
	}
}
