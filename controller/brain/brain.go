package brain

import (
	"math"
	"neuroSnake/controller"
	"neuroSnake/controller/brain/neuron"
)

/*
brain is another runner which
calls operations on snake
*/

type brain struct {
	inputDendrits    []neuron.Dendrit
	resultNeurons    []neuron.Neuron
	validateFunction func() []bool
}

func simpleValidator(impulces []float64) []bool {
	maxi, max := 0, impulces[0]
	validatorResults := make([]bool, len(impulces))
	for i := range impulces {
		if impulces[i] > max {
			maxi, max = i, impulces[i]
		}
	}

	for i := range validatorResults {
		if i == maxi {
			validatorResults[i] = true
		} else {
			validatorResults[i] = false
		}
	}

	return validatorResults

}

func NewBrain(brainMaker func() ([]neuron.Dendrit, []neuron.Neuron), validateFunction func([]float64) []bool) controller.Runner {
	inputs, outputs := brainMaker()
	return &brain{inputs, outputs, validateFunction}
}

func (b *brain) Run() chan error {
	return nil
}
