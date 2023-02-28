package test

import (
	"graph_robot/database"
	"graph_robot/graph_structure"
	"testing"
)

func TestCreateOne(t *testing.T) {
	neure := graph_structure.Neure{
		AxonSynapse: graph_structure.Synapse{
			NextNeureID: 1,
			Weight:      222,
		},
		// DendritesLinkNum:       3,
		NeureType:              true,
		ElectricalConductivity: 443,
	}
	neure.CreateNeureInDB()

	if neure.ThisNeureId == 0 {
		t.Error("id is 0")
	}

	t.Logf("Success ####%+v", neure)
}

func TestCreateMulti(t *testing.T) {
	neures := graph_structure.CreateNewNeures(10)
	if len(neures) == 0 {
		t.Error("nothing created")
	}
}

func TestConnectNextNuere(t *testing.T) {
	neure := graph_structure.Neure{}
	neure.GetNeureFromDatabaseById(2)
	neure.ConnectNextNuere(1)

	neure = graph_structure.Neure{}
	neure.GetNeureFromDatabaseById(2)
	if neure.ThisNeureId != 2 {
		t.Error("this id is wrong id")
	}
	if neure.AxonSynapse.NextNeureID != 1 {
		t.Error("Link fail")
	}

	dbModel := database.NeureData{}
	dbModel.GetNeureDataById(2)
	if dbModel.Linked != true {
		t.Error("db update linked fail")
	}
}

func TestLoadNeure(t *testing.T) {
	neure := graph_structure.Neure{}
	neure.GetNeureFromDatabaseById(1)

	if neure.ThisNeureId != 1 {
		t.Error("get fail")
	}
}

func TestGetUnlinked(t *testing.T) {
	amount := 3
	neures := database.GetUnlinkedNeures(amount)
	if len(neures) != 3 || neures[0].Linked != false {
		t.Error("not get enough unlinked or linked is not false")
	}
}

// notice: Network was abandoned

// func TestLoad(t *testing.T) {
// 	firstLink := []int64{2}
// 	network := graph_structure.NetWork{
// 		Neures: make(map[int64]*graph_structure.Neure),
// 	}
// 	network.LoadNetwork(firstLink[0])
// 	if len(network.Neures) < 1 {
// 		t.Error("no data")
// 	}
// 	t.Logf("###%+v", network.Neures)
// 	t.Log("$$$", network.NeureOrder)
// }

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

// func TestToByte(t *testing.T) {
// 	n := &Neure
// 	js, err := json.Marshal(*n)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log("length of json:", len(js))
// 	t.Log("js:", string(js))
// 	t.Log("length of string:", len(string(js)))

// 	by := []byte{123, 34, 97, 115, 34, 58, 123, 34, 110, 110, 34, 58, 52, 51, 50, 52, 51, 50, 52, 50, 52, 44, 34, 119, 119, 34, 58, 52, 51, 53, 52, 51, 53, 125, 44, 34, 100, 108, 34, 58, 51, 51, 51, 52, 52, 44, 34, 110, 108, 100, 34, 58, 52, 51, 52, 51, 44, 34, 110, 116, 34, 58, 116, 114, 117, 101, 44, 34, 101, 108, 99, 34, 58, 52, 52, 50, 51, 52, 50, 51, 125}
// 	var aaa graph_structure.Neure
// 	err = json.Unmarshal(by, &aaa)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log("aaa.Weight", aaa.AxonSynapse.Weight)

// 	b := graph_structure.Struct2Byte(&Neure)
// 	t.Log("length of byte:", len(b))

// 	_ = graph_structure.Byte2Struct(b)
// }
