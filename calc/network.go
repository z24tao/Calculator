package calc

import (
	"fmt"
)

type Network struct {
	Inputs  []*float64
	Neurons [][]*Neuron
}

type Design struct {
	InputSize int
	Layers    []int
}

func (n *Network) Train(inputs []float64, answers []float64, learningRate float64) {
	if len(n.Inputs) != len(inputs) {
		panic("input lengths do not match")
	}
	outputLayer := n.Neurons[len(n.Neurons)-1]
	if len(outputLayer) != len(answers) {
		panic("output lengths do not match")
	}

	fmt.Println("inputs:", inputs)
	n.setInput(inputs)
	n.forwardPropagate()
	n.backwardPropagate(answers, learningRate)

	results := make([]float64, 0)
	for _, outputNode := range outputLayer {
		results = append(results, *outputNode.Result)
	}

	fmt.Println("results: ", results)
	fmt.Println("expected:", answers)
	fmt.Println("error:   ", meanSquaredError(results, answers))
	fmt.Println()
}

func (n *Network) setInput(values []float64) {
	for i := 0; i < len(n.Inputs); i++ {
		*(n.Inputs[i]) = values[i]
	}
}

func (n *Network) forwardPropagate() {
	for _, layer := range n.Neurons {
		for _, node := range layer {
			node.Compute()
		}
	}
}

func (n *Network) backwardPropagate(answers []float64, learningRate float64) {
	errorTerms := n.calculateErrorTerms(answers)
	n.updateWeights(errorTerms, learningRate)
}

func (n *Network) calculateErrorTerms(answers []float64) [][]float64 {
	outputLayer := n.Neurons[len(n.Neurons)-1]
	if len(outputLayer) != len(answers) {
		panic("output lengths do not match")
	}

	//calculate error terms backwards
	errorTerms := make([][]float64, len(n.Neurons))

	//output layer
	outputLayerError := make([]float64, len(outputLayer))
	for i, outputNode := range outputLayer {
		outputLayerError[i] = (*outputNode.Result - answers[i]) * outputNode.Activation.Derivative(outputNode.WeightedSum)
	}
	errorTerms[len(n.Neurons)-1] = outputLayerError

	//previous layers
	for i := len(n.Neurons) - 2; i >= 0; i-- {
		curLayer := n.Neurons[i]
		nextLayer := n.Neurons[i+1]
		curLayerError := make([]float64, len(curLayer))
		//TODO RIGHT NOW WE TRUST THAT ALL INDICES MATCH, THIS IS VERY FRAGILE
		for j, curLayerNode := range curLayer {
			curLayerError[j] = float64(0)
			for k, nextLayerNode := range nextLayer {
				deltaIK := errorTerms[i+1][k]
				weightJK := nextLayerNode.Inputs[j].Weight
				activationDerivative := curLayerNode.Activation.Derivative(curLayerNode.WeightedSum)
				curLayerError[j] += deltaIK * weightJK * activationDerivative
			}
		}
		errorTerms[i] = curLayerError
	}

	return errorTerms
}

func (n *Network) updateWeights(errorTerms [][]float64, learningRate float64) {
	for i, curLayer := range n.Neurons {
		for j, curLayerNode := range curLayer {
			valueSum := float64(0)
			for _, synapse := range curLayerNode.Inputs {
				fmt.Print("updating weight from ", synapse.Weight, " to ")
				synapse.Weight -= learningRate * errorTerms[i][j] * *(synapse.Value)
				fmt.Println(synapse.Weight)
				valueSum += *(synapse.Value)
			}
			fmt.Print("updating bias from ", curLayerNode.Bias, " to ")
			curLayerNode.Bias -= learningRate * errorTerms[i][j]
			fmt.Println(curLayerNode.Bias)
		}
	}
}
