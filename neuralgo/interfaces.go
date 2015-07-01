package neuralgo

type ActivationFunction func(float64) float64

type NeuronLayer interface {
	ProcessInput([]float64) []float64
}
