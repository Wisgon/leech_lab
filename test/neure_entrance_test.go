package test

import (
	"graph_robot/config"
	"graph_robot/neure"
	"testing"
)

func TestWrite2File(t *testing.T) {
	ne := neure.NeureEntrance{
		EntranceType: config.EntranceTypes["eyes"],
		NeuresIds:    []int64{1, 2, 3},
	}
	ne.Save2File()
}

func TestReadFile(t *testing.T) {
	ne := neure.NeureEntrance{
		EntranceType: config.EntranceTypes["eyes"],
	}
	ne.LoadFromFile()
	if ne.NeuresIds[0] != 1 {
		t.Error("not work")
	}
}
