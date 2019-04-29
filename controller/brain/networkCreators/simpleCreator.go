package networkCreators

import (
	"math"
	"math/rand"
	"neuroSnake/controller/brain/neuron"
	"time"
)

func Sigmoid(x float64, isDerive bool) float64 {
	if isDerive {
		return x * (1 - x)
	}
	return 1 / (1 + math.Exp(x))
}

/*
	takes inputsNum inputs and moves them into
	outputsNum outputs using layersNum slices between
*/
func simpleNetwork(inputsNum uint8, outputsNum uint8, layersNum uint8) ([]neuron.Dendrit, []neuron.Neuron) {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	var prevNeurons []neuron.Neuron = nil
	var initialInputs []neuron.Dendrit = nil

	layerDiff := uint8(math.Round(float64(inputsNum-outputsNum) / float64(outputsNum)))
	currentLayer := inputsNum

	for j := uint8(0); j < layersNum; j++ {
		inputs := make([]neuron.Dendrit, currentLayer)
		if initialInputs == nil {
			initialInputs = inputs
		}

		currentLayer -= layerDiff
		if currentLayer < outputsNum {
			currentLayer = outputsNum
		}

		neurons := make([]neuron.Neuron, currentLayer)
		// creating dendrits with random weights
		for i := range inputs {
			randFloat := rnd.Float64()
			inputs[i] = neuron.NewDendrit(randFloat - math.Round(randFloat))
		}

		// making this dendrits axons to previouse neuron layer
		for i := range prevNeurons {
			prevNeurons[i].SetAxon(inputs[i])
		}

		connectNum := 3
		connectMove := int(math.Trunc(float64(currentLayer+layerDiff) / float64(currentLayer)))

		// making dendrits connected to other neurons
		for i := 0; i < len(inputs); i += connectMove {
			neurons[i/connectMove] = neuron.NewNeuron(inputs[i:i+connectNum], nil, Sigmoid)
		}
		prevNeurons = neurons

		// start neuron layer
		for i := range neurons {
			neurons[i].Activate()
		}
	}

	return initialInputs, prevNeurons
}
