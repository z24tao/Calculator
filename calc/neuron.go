package calc

type Synapse struct {
	Weight float64
	Value  *float64
}

type Neuron struct {
	Inputs      []*Synapse
	Bias        float64
	WeightedSum float64
	Activation  ActivationFunction
	Result      *float64
}

func (n *Neuron) Compute() {
	n.WeightedSum = n.Bias
	for _, input := range n.Inputs {
		n.WeightedSum += *input.Value * input.Weight
	}
	*n.Result = n.Activation.Body(n.WeightedSum)
}
