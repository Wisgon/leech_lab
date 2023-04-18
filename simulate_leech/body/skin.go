package body

import (
	"graph_robot/neure"
	"graph_robot/simulate_leech/utils"
)

type Skin struct {
	Position string                    // position of this part of Skin
	Entrance map[string][]*neure.Neure // signal entrance
}

func (s *Skin) InitSkin() {
	// get entrance of this position of this part of skin
}

func (s *Skin) Temperature2NeuralSignal(temperature float64) {
	// transform temperature value get from env to neural signal
	switch {
	case temperature > 12 && temperature < 45:
		// normal temp, activate normal temp neure
		for _, n := range s.Entrance["skin_entrance_normalTemperature"] {
			go utils.SignalPass(n)
		}
		// todo: next entrance
	}
}
