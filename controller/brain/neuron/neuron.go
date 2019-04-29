package neuron

type (
	neuron struct {
		dendrits  []Dendrit
		axon      Dendrit
		normalise func(float64, bool) float64
		lastRes   *float64
	}

	Neuron interface {
		Activate() chan float64
		LastRes() float64
		SetAxon(Dendrit)
	}
)

func NewNeuron(dendrits []Dendrit, axon Dendrit, normFunc func(float64, bool) float64) Neuron {
	return &neuron{dendrits, axon, normFunc, nil}
}

func (n *neuron) Activate() chan float64 {
	result := make(chan float64, 1)

	go func() {
		inputs := make(chan float64, len(n.dendrits))
		for i := range n.dendrits {
			n.dendrits[i].Derive(inputs)
		}

		summ := float64(0)
		for i := 0; i < len(n.dendrits); i++ {
			summ += <-inputs
		}

		summ = n.normalise(summ, false)

		result <- summ
		if n.axon != nil {
			n.axon.SetImpulse(summ)
		}
		n.lastRes = &summ
	}()

	return result
}

/*
 no guarantee that the returned value will be actually the last
*/
func (n *neuron) LastRes() float64 {
	return *n.lastRes
}

func (n *neuron) SetAxon(axon Dendrit) {
	n.axon = axon
}
