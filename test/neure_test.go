package test

import (
	"graph_robot/neure"
	"testing"
)

func TestNeureLink(t *testing.T) {
	neure1 := neure.Neure{}
	neure2 := neure.Neure{}
	neure1.CreateNeureInDB("testing_neure")
	neure2.CreateNeureInDB("testing_neure")
	neure3 := neure.Neure{}
	neure3.CreateNeureInDB("testing_neure")

	neure1.ConnectNextNuere(&neure2)
	neure2.ConnectNextNuere(&neure3)

	key1, key2, key3 := neure1.ThisNeureId, neure2.ThisNeureId, neure3.ThisNeureId
	t.Log("key2:", key2)

	neure1 = neure.Neure{}
	neure2 = neure.Neure{}
	neure3 = neure.Neure{}

	neure1.GetNeureFromDbById(key1)
	neure2.GetNeureFromDbById(key2)
	neure3.GetNeureFromDbById(key3)

	if len(neure1.Synapses) != 0 && neure1.Synapses[0].GetNextId() == key2 && len(neure2.NowLinkedDendritesIds) != 0 && neure2.NowLinkedDendritesIds[0] == key1 && len(neure3.NowLinkedDendritesIds) != 0 && neure3.NowLinkedDendritesIds[0] == key2 {
		t.Logf("Success~~~~n1:%+v, n2:%+v", neure1, neure2)
	} else {
		t.Error("Link Fail")
	}
}

func TestDeleteNeure(t *testing.T) {
	key2 := "testing_neure@1"
	key1 := "testing_neure@0"
	key3 := "testing_neure@2"

	neure2 := neure.Neure{}
	neure2.GetNeureFromDbById(key2)
	neure2.DeleteNeure()

	neure1 := neure.Neure{}
	neure1.GetNeureFromDbById(key1)
	neure3 := neure.Neure{}
	neure3.GetNeureFromDbById(key3)

	if len(neure1.Synapses) == 0 && len(neure3.NowLinkedDendritesIds) == 0 {
		t.Log("Success~~~~~:)")
	} else {
		t.Error("Fail!!!:(")
	}
}
