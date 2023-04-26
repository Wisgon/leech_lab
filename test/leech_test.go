package test

import (
	"graph_robot/database"
	leech "graph_robot/simulate_leech"
	"testing"
)

func TestInitLeechCreateNeure(t *testing.T) {
	leechObj := leech.Leech{}
	leechObj.InitLeech()

	allNeuresBytes := database.ValueAllDbScan(func(result []byte) bool {
		// raw prefix will set an empty value in db, if get in Byte2Struct, will panic
		return len(result) != 8
	}, false)
	if len(*allNeuresBytes) == 0 {
		t.Errorf("get all fail...\n")
	} else {
		for _, v := range *allNeuresBytes {
			t.Logf("value: %s\n", string(v))
		}
	}
	t.Log("Success~~~")
}
