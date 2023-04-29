package body

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"sync"
	"time"
)

type Muscle struct {
	mu            sync.Mutex
	MoveDirection string   `json:"a"`
	Neures        []string `json:"b"`
	KeyPrefix     string   `json:"d"`
}

func (m *Muscle) InitMuscle(wg *sync.WaitGroup) {
	defer wg.Done()
	m.createNeures()

	// finally save this to database with keyPrefix+config.PrefixNumSplitSymbol+"collection"
	dataByte := m.struct2Byte()
	database.CreateData(dataByte, m.KeyPrefix+config.PrefixNumSplitSymbol+"collection")
}

func (m *Muscle) createNeures() {
	neureObj := neure.CreateOneNeure(m.KeyPrefix, &neure.Neure{
		NeureType:              "normal",
		LastTimeActivate:       time.Now(),
		LastTimeResetNowWeight: time.Now(),
	})
	m.Neures = append(m.Neures, neureObj.ThisNeureId)
}

func (m *Muscle) struct2Byte() []byte {
	dataByte, err := json.Marshal(m)
	if err != nil {
		panic("json marshal error: " + err.Error())
	}
	return dataByte
}
