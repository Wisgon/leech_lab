package test

import (
	"graph_robot/config"
	"graph_robot/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	database.InitDb(config.TestDataPath, config.SeqBandwidth)
	code := m.Run()
	database.CloseDb()
	os.Exit(code)
}
