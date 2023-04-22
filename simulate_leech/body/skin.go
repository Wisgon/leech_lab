package body

type Skin struct {
	Position string   // position of this part of Skin
	Neures   []string // neures of this position
}

func (s *Skin) InitSkin() {
	// get entrance of this position of this part of skin
}

func (s *Skin) Temperature2NeuralSignal(temperature float64) {
	// transform temperature value get from env to neural signal
	switch {
	case temperature > 12 && temperature < 45:

	}
}
