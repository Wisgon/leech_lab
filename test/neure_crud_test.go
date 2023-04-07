package test

import (
	"fmt"
	"testing"

	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
)

func TestCreateOne(t *testing.T) {
	neureIns := neure.Neure{
		ElectricalConductivity: 443,
	}
	key := database.GetKeyFromPrefix("testing_neure")
	neureIns.ThisNeureId = key
	database.CreateNeure(neureIns.Struct2Byte(), key)

	t.Logf("Success ####%+v", neureIns)
}

func TestGetNeure(t *testing.T) {
	key := "testing_neure@0"
	neureObj := neure.Neure{}
	neureObj.GetNeureFromDbById(key)
	if neureObj.ThisNeureId != key {
		t.Error("get wrong data")
	}
	t.Logf("Success~~, %+v\n", neureObj)
}

func TestUpdateNeure(t *testing.T) {
	key := "testing_neure@1"
	neureByte := database.GetNeure(key)
	neureObj := neure.Neure{}
	neureObj.Byte2Struct(neureByte)

	neureObj.ElectricalConductivity = 111
	database.UpdateNeure(neureObj.Struct2Byte(), neureObj.ThisNeureId)

	neureByte = database.GetNeure(key)
	neureObj = neure.Neure{}
	neureObj.Byte2Struct(neureByte)
	if neureObj.ElectricalConductivity != 111 {
		t.Error("update fail")
	}
	t.Log("Success ~~~")
}

func TestDelete(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			errMsg := fmt.Sprint(err)
			if errMsg == "Key not found" {
				t.Log("delete success~~~")
			}
		}
	}()
	key := "testing_neure@2"
	database.DeleteNeure(key)

	_ = database.GetNeure(key)
}

func TestScanAllKey(t *testing.T) {
	database.CheckAllKey()
	t.Log("success~~~")
}

func TestScanAll(t *testing.T) {
	allNeuresBytes := database.ValueAllDbScan(func(result []byte) bool {
		n := neure.Neure{}
		if len(result) == 8 {
			// raw prefix will set an empty value in db, if get in Byte2Struct, will panic
			return false
		}
		n.Byte2Struct(result)
		return n.ElectricalConductivity == 443
	}, false)
	if len(*allNeuresBytes) == 0 {
		t.Errorf("get all fail...\n")
	} else {
		allNeures := neure.TurnNeureBytes2Neures(allNeuresBytes)
		for _, v := range *allNeures {
			t.Logf("key: %s", v.ThisNeureId)
			if v.ElectricalConductivity != 443 {
				t.Errorf("get wrong neure, filter not work, value is : %d", v.ElectricalConductivity)
			}
		}
		t.Logf("neures are: %d\n", len(*allNeures))
	}

	// test first flag
	firstNeuresBytes := database.ValueAllDbScan(func(result []byte) bool {
		n := neure.Neure{}
		if len(result) == 8 {
			// raw prefix will set an empty value in db, if get in Byte2Struct, will panic
			return false
		}
		n.Byte2Struct(result)
		return n.ElectricalConductivity == 443
	}, true)
	if len(*firstNeuresBytes) != 1 {
		t.Error("len not equal 1")
	} else {
		t.Log("Ok getting first")
	}
}

func TestScanPrefixAll(t *testing.T) {
	prefixedNeureBytes := database.ValueAndPrefixScan("testing_neure"+config.PrefixNumSplitSymbol, func(result []byte) bool {
		n := neure.Neure{}
		n.Byte2Struct(result)
		return n.ElectricalConductivity == 443
	}, false)
	prefixedNeures := neure.TurnNeureBytes2Neures(prefixedNeureBytes)
	if len(*prefixedNeures) == 0 {
		t.Error("no data found")
	} else {
		for k := range *prefixedNeures {
			t.Logf("key: %s", (*prefixedNeures)[k].ThisNeureId)
			if (*prefixedNeures)[k].ElectricalConductivity != 443 {
				t.Error("wrong data found")
			}
		}
		t.Logf("neures are: %d", len(*prefixedNeures))
	}

	// test first flag
	firstPrefixedNeureBytes := database.ValueAndPrefixScan("testing_neure"+config.PrefixNumSplitSymbol, func(result []byte) bool {
		n := neure.Neure{}
		n.Byte2Struct(result)
		return n.ElectricalConductivity == 443
	}, true)
	if len(*firstPrefixedNeureBytes) != 1 {
		t.Error("not found 1")
	} else {
		t.Logf("neures are: %d", len(*firstPrefixedNeureBytes))
	}
}

func TestKeyOnlyPrefixScan(t *testing.T) {
	neureBytesKeys := database.KeyOnlyPrefixScan("testing_neure" + config.PrefixNumSplitSymbol)
	if len(*neureBytesKeys) == 0 {
		t.Error("no data found")
	}
	for _, v := range *neureBytesKeys {
		t.Logf("key: %s", v)
	}
}

func TestManuallyCreate(t *testing.T) {
	mtxn := database.ManuallyTransaction{}
	defer mtxn.Close()

	mtxn.Init()
	var neures = make(map[string]*neure.Neure)
	for i := 0; i < 3012; i++ {
		key := database.GetKeyFromPrefix("testing_neure")
		n := neure.Neure{
			ElectricalConductivity: 109090,
			ThisNeureId:            key,
		}
		neures[key] = &n
		mtxn.Create(key, n.Struct2Byte())
	}

	t.Logf("success~~~, len: %d", len(neures))
}

func TestManuallyUpdate(t *testing.T) {
	mtxn := database.ManuallyTransaction{}
	defer mtxn.Close()
	mtxn.Init()

	neures := make(map[string]*neure.Neure)
	keys := []string{}

	neureBytesKeys := database.KeyOnlyPrefixScan("testing_neure" + config.PrefixNumSplitSymbol)
	if len(*neureBytesKeys) == 0 {
		t.Error("no data found")
	}
	for _, v := range *neureBytesKeys {
		keys = append(keys, string(v))
	}

	for _, key := range keys {
		neureByte := database.GetNeure(key)
		neureObj := neure.Neure{}
		neureObj.Byte2Struct(neureByte)
		neures[key] = &neureObj
	}

	for key, n := range neures {
		n.ElectricalConductivity = 9001
		mtxn.Update(key, n.Struct2Byte())
	}

	neureByte := database.GetNeure("testing_neure@1")
	neureObj := neure.Neure{}
	neureObj.Byte2Struct(neureByte)
	if neureObj.ElectricalConductivity != 9001 {
		t.Errorf("Update fail: ele:%d\n", neureObj.ElectricalConductivity)
	}

	t.Logf("success~~~, len: %d", len(neures))
}

func TestManuallyDelete(t *testing.T) {
	mtxn := database.ManuallyTransaction{}
	defer mtxn.Close()
	mtxn.Init()

	keys := []string{}
	neureBytesKeys := database.KeyOnlyPrefixScan("testing_neure" + config.PrefixNumSplitSymbol)
	if len(*neureBytesKeys) == 0 {
		t.Error("no data found")
	}
	for _, v := range *neureBytesKeys {
		keys = append(keys, string(v))
	}
	for _, key := range keys {
		mtxn.Delete(key)
	}

	t.Log("Success~~~")
}
