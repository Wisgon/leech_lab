package test

import (
	"encoding/json"
	"graph_robot/neure"
	"sync"
	"testing"
)

func TestNeureLink(t *testing.T) {
	neure1 := neure.CreateOneNeure("testing_neure", &neure.Neure{
		NeureType: "normal",
	})
	neure2 := neure.CreateOneNeure("testing_neure", &neure.Neure{
		NeureType: "normal",
	})
	neure3 := neure.CreateOneNeure("testing_neure", &neure.Neure{
		NeureType: "normal",
	})

	t.Log("neure2 thisid:", neure2.ThisNeureId)
	neure1.ConnectNextNuere(neure2.ThisNeureId)
	neure2.ConnectNextNuere(neure3.ThisNeureId)

	t.Logf("n1: %+v##########\n n2: %+v##########----\n n3: %+v##########--\n", neure1, neure2, neure3)

	key1, key2, _ := neure1.ThisNeureId, neure2.ThisNeureId, neure3.ThisNeureId
	t.Log("key2:", key2)

	if len(neure1.Synapses) != 0 && neure1.Synapses[0].NextNeureID == key2 && len(neure2.NowLinkedDendritesIds) != 0 && neure2.NowLinkedDendritesIds[0] == key1 && len(neure3.NowLinkedDendritesIds) != 0 && neure3.NowLinkedDendritesIds[0] == key2 {
		t.Logf("Success~~~~n1:%+v, n2:%+v", neure1, neure2)
	} else {
		t.Error("Link Fail")
	}
}

func TestDeleteNeure(t *testing.T) {
	key2 := "testing_neure@1"
	key1 := "testing_neure@0"
	key3 := "testing_neure@2"

	neure2 := neure.GetNeureById(key2)
	neure.DeleteNeure(neure2)
	neure1 := neure.GetNeureById(key1)
	neure3 := neure.GetNeureById(key3)

	n1, _ := json.Marshal(neure1.Struct2Byte())
	t.Log("bytes!!!!!!!!!", string(n1))

	if len(neure1.Synapses) == 0 && len(neure3.NowLinkedDendritesIds) == 0 {
		t.Log("Success~~~~~:)")
	} else {
		t.Error("Fail!!!:(")
	}
}

func TestNeureMap(t *testing.T) {
	// test concurrently increase ElectricalConductivity
	neure1 := neure.CreateOneNeure("testing_neure", &neure.Neure{
		ElectricalConductivity: 0,
	})
	key := neure1.ThisNeureId
	neure1 = neure.GetNeureById(key)

	var wg sync.WaitGroup

	wg.Add(3)

	addNeureElectrical := func(n *neure.Neure) {
		for i := 0; i < 10000; i++ {
			n.ChangeElectricalConductivity(1, "add")
		}
		wg.Done()
	}

	go addNeureElectrical(neure1)
	go addNeureElectrical(neure1)
	go addNeureElectrical(neure1)

	wg.Wait()

	if neure1.ElectricalConductivity != 30000 {
		t.Error("Fail:", neure1.ElectricalConductivity)
	} else {
		t.Log("success~~~")
	}
}
