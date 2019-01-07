package main

import (
	"calculator/calc"
	"calculator/data"
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(1)
	network := calc.CreateNetwork(calc.Design{
		InputSize: 2,
		Layers:    []int{1},
	})

	trainLength := 1000000
	//learningRate := 0.00000001
	learningRate := 0.00001
	for i := 0; i < trainLength; i++ {
		a, b, c := data.AdditionRandom(-100, 100)
		network.Train([]float64{a, b}, []float64{c}, float64(learningRate))
	}

	for _, layer := range network.Neurons {
		for _, node := range layer {
			fmt.Print("weights: ")
			for _, synapse := range node.Inputs {
				fmt.Print(synapse.Weight, " ")
			}
			fmt.Println()
			fmt.Println("bias:", node.Bias)
		}
		fmt.Println()
	}
}
