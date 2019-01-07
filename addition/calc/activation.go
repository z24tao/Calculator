package calc

type ActivationFunction struct {
	Body       func(float64) float64
	Derivative func(float64) float64
}

var identity = ActivationFunction{
	Body: func(in float64) float64 {
		return in
	},
	Derivative: func(in float64) float64 {
		return 1
	},
}
