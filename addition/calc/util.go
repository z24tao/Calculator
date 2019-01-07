package calc

import "math/rand"

func meanSquaredError(left, right []float64) float64 {
	if len(left) != len(right) {
		panic("left and right lengths do not match")
	}

	result := float64(0)
	for i := 0; i < len(left); i++ {
		result += (left[i] - right[i]) * (left[i] - right[i])
	}

	return result
}

func connect(source *float64, next *Neuron, weight float64) {
	synapse := &Synapse{
		Weight: weight,
		Value:  source,
	}

	next.Inputs = append(next.Inputs, synapse)
}

func CreateNetwork(design Design) *Network {
	network := &Network{}

	//create inputs
	if design.InputSize <= 0 {
		panic("input size must be at least 1")
	}
	network.Inputs = make([]*float64, design.InputSize)
	for i := 0; i < design.InputSize; i++ {
		var input float64
		network.Inputs[i] = &input
	}

	//create neurons
	if len(design.Layers) == 0 {
		panic("network must have at least one layer")
	}
	network.Neurons = make([][]*Neuron, len(design.Layers))
	for i, size := range design.Layers {
		if size <= 0 {
			panic("layer size must be at least 1")
		}
		layer := make([]*Neuron, size)
		for j := 0; j < size; j++ {
			var nodeResult float64
			node := &Neuron{
				Inputs:     make([]*Synapse, 0),
				Bias:       rand.Float64() - 0.5,
				//Bias:       0,
				Activation: identity,
				Result:     &nodeResult,
			}
			layer[j] = node
		}
		network.Neurons[i] = layer
	}

	for _, input := range network.Inputs {
		for _, neuron := range network.Neurons[0] {
			connect(input, neuron, rand.Float64()-0.5)
			//connect(input, neuron, 1)
		}
	}

	//connect - remaining layers
	for inputLayerIndex := 0; inputLayerIndex < len(network.Neurons)-1; inputLayerIndex++ {
		inputLayer := network.Neurons[inputLayerIndex]
		outputLayer := network.Neurons[inputLayerIndex+1]

		for _, inputNode := range inputLayer {
			for _, outputNode := range outputLayer {
				connect(inputNode.Result, outputNode, rand.Float64()-0.5)
				//connect(inputNode.Result, outputNode, 1)
			}
		}
	}

	return network
}
