package test

import (
	"encoding/json"
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
		// DendritesLinkNum:       3,
		NeureType:              true,
		ElectricalConductivity: 443,
	}
	neureByte := database.NeureDb{
		Neure: graph_structure.Struct2Byte(&neure),
	}
	id := database.SaveNeure(neureByte)

	if id == 0 {
		t.Error("id is 0")
	}

	t.Log("Success ####", id)
}

func TestLoad(t *testing.T) {
	firstLink := []int64{2}
	network := graph_structure.NetWork{
		Neures: make(map[int64]*graph_structure.Neure),
	}
	network.LoadNetwork(firstLink[0])
	if len(network.Neures) < 1 {
		t.Error("no data")
	}
	t.Logf("###%+v", network.Neures)
	t.Log("$$$", network.NeureOrder)
}

// notice: Network was abandoned

// func TestNetworkUpdate(t *testing.T) {
// 	firstLink := []int64{4}
// 	network := graph_structure.NetWork{
// 		Neures: make(map[int64]*graph_structure.Neure),
// 	}
// 	network.LoadNetwork(firstLink[0])
// 	t.Logf("$$$$%d", network.Neures[4].AxonSynapse.Weight)
// 	network.Neures[4].AxonSynapse.Weight = 777
// 	network.NeedUpdateNeures = append(network.NeedUpdateNeures, firstLink[0])
// 	network.SaveNetwork()

// 	network2 := graph_structure.NetWork{
// 		Neures: make(map[int64]*graph_structure.Neure),
// 	}
// 	network2.LoadNetwork(firstLink[0])
// 	if network2.Neures[4].AxonSynapse.Weight != 777 {
// 		t.Error("update fail")
// 	}
// 	t.Logf("###%+d", network2.Neures[4].AxonSynapse.Weight)

// }

func TestToByte(t *testing.T) {
	n := &Neure
	js, err := json.Marshal(*n)
	if err != nil {
		t.Error(err)
	}
	t.Log("length of json:", len(js))
	t.Log("js:", string(js))
	t.Log("length of string:", len(string(js)))

	by := []byte{123, 34, 97, 115, 34, 58, 123, 34, 110, 110, 34, 58, 52, 51, 50, 52, 51, 50, 52, 50, 52, 44, 34, 119, 119, 34, 58, 52, 51, 53, 52, 51, 53, 125, 44, 34, 100, 108, 34, 58, 51, 51, 51, 52, 52, 44, 34, 110, 108, 100, 34, 58, 52, 51, 52, 51, 44, 34, 110, 116, 34, 58, 116, 114, 117, 101, 44, 34, 101, 108, 99, 34, 58, 52, 52, 50, 51, 52, 50, 51, 125}
	var aaa graph_structure.Neure
	err = json.Unmarshal(by, &aaa)
	if err != nil {
		t.Error(err)
	}
	t.Log("aaa.Weight", aaa.AxonSynapse.Weight)

	b := graph_structure.Struct2Byte(&Neure)
	t.Log("length of byte:", len(b))

	_ = graph_structure.Byte2Struct(b)
}

func TestGetUnlinked(t *testing.T) {
	amount := 3
	neures := database.GetUnlinkedNeures(amount)
	t.Log("length of neures:", len(neures))
}

func TestUpdateLinked(t *testing.T) {
	neureDb := database.GetNeureById(1)
	database.UpdateLinked(neureDb.ID)
}
