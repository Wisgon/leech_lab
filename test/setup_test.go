package test

import (
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"os"
	"testing"
)

func cleanup() {
	// save neure map
	neure.NeureMap.Range(func(key, value any) bool {
		neureObj := value.(*neure.Neure)
		neureObj.UpdateNeure2DB()
		return true
	})
	fmt.Println("closing db~~~")
	database.CloseDb()
	// some other cleanup here ~~~
}

func TestMain(m *testing.M) {
	defer func() {
		if r := recover(); r != nil {
			cleanup()
		}
	}()
	database.InitDb(config.TestDataPath, config.SeqBandwidth)
	code := m.Run()
	cleanup()
	os.Exit(code)
}
