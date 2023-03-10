package test

import (
	"encoding/json"
	"graph_robot/neure"
	"testing"
)

var Neure = neure.Neure{
	AxonSynapse: neure.Synapse{
		NextNeureID: 432432424,
		Weight:      435435,
	},
	// DendritesLinkNum:       33344,
	NowLinkedDendritesNum:  4343,
	NeureType:              true,
	ElectricalConductivity: 4423423,
}

func BenchmarkJson(*testing.B) {
	// use json is more faster
	for i := 0; i < 1000; i++ {
		nb, _ := json.Marshal(Neure)
		_ = string(nb)
		var neu neure.Neure
		_ = json.Unmarshal(nb, &neu)
	}

}

// func BenchmarkSb(*testing.B) {
// 	for i := 0; i < 1000; i++ {
// 		sb := neure.Struct2Byte(&Neure)
// 		_ = neure.Byte2Struct(sb)
// 	}
// }
