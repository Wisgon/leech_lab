package test

import (
	"graph_robot/database"
	"graph_robot/graph_structure"
	"testing"
)

func TestCreate(t *testing.T) {
	neure := graph_structure.Neure{
		AxonSynapse: graph_structure.Synapse{
			NextNeureID: 1,
			Weight:      222,
		},
		DendritesLinkNum:       3,
		NeureType:              true,
		ElectricalConductivity: 443,
	}
	neureByte := database.Neures{
		Neure: graph_structure.Struct2Byte(&neure),
	}
	id := database.SaveNeure(neureByte)

	if id == 0 {
		t.Error("id is 0")
	}

	t.Log("Success ####", id)
}
