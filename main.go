package main

import (
	"./neuralgo"
	"fmt"
)

func main() {

	neuralNetwork := neuralgo.NewNeuralNetwork(
		2, // num inputs
		1, // num outputs
		5, // num hidden
		3, // num neurons/hidden
		neuralgo.Sigmoid)
	for i := 0; i < 100; i++ {
		fmt.Printf("%f\n", neuralNetwork.ProcessInput([]float64{float64(float64(i) / 100), float64(float64(i) / 100)}))
	}
}
