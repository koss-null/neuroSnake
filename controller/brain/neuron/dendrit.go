package neuron

type (
	dendrit struct {
		impulse chan float64
		weight  float64
	}

	Dendrit interface {
		SetImpulse(impulse float64)
		Derive(result chan<- float64)
		Change(value float64)
	}
)

func NewDendrit(initWeight float64) Dendrit {
	return &dendrit{make(chan float64, 1), initWeight}
}

func (d *dendrit) SetImpulse(impulse float64) {
	d.impulse <- impulse
}

func (d *dendrit) Derive(result chan<- float64) {
	go func() {
		result <- d.weight * <-d.impulse
	}()
}

func (d *dendrit) Change(value float64) {
	d.weight = value
}

func (d *dendrit) Weight() float64 {
	return d.weight
}
