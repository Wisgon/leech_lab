package test

import (
	"encoding/json"
	"graph_robot/graph_structure"
	"testing"
)

var Neure = graph_structure.Neure{
	AxonSynapse: graph_structure.Synapse{
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
		var neu graph_structure.Neure
		_ = json.Unmarshal(nb, &neu)
	}

}

// func BenchmarkSb(*testing.B) {
// 	for i := 0; i < 1000; i++ {
// 		sb := graph_structure.Struct2Byte(&Neure)
// 		_ = graph_structure.Byte2Struct(sb)
// 	}
// }
