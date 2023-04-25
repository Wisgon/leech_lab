package body

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"strings"
	"sync"
	"time"
)

type Skin struct {
	mu            sync.Mutex
	Position      string   `json:"a"` // position of this part of Skin
	SkinNeureType string   `json:"b"` // type of skin neure
	Neures        []string `json:"c"` // neure id of this position
	KeyPrefix     string   `json:"d"`
}

func (s *Skin) InitSkin(wg *sync.WaitGroup, processController *sync.Map) {
	defer wg.Done()
	s.createNeures()

	// finally save this to database with keyPrefix+config.PrefixNumSplitSymbol+"collection"
	dataByte := s.struct2Byte()
	database.CreateData(dataByte, s.KeyPrefix+config.PrefixNumSplitSymbol+"collection")
}

func (s *Skin) Temperature2NeuralSignal(temperature float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// transform temperature value get from env to neural signal
	switch {
	case temperature > 12 && temperature < 45:

	}
}

func (s *Skin) createNeures() {
	var neureNum int
	if strings.Contains(s.SkinNeureType, "normal") {
		neureNum = config.EachSkinPositionSurfaceNeureNum
	} else if strings.Contains(s.SkinNeureType, "bigger") {
		neureNum = config.EachSkinPositionDeeperNeureNum
	} else if strings.Contains(s.SkinNeureType, "extremely") {
		neureNum = config.EachSkinPositionDeepestNeureNum
	} else {
		panic("Unknow skin neure type")
	}
	for i := 0; i < neureNum; i++ {
		neureObj := neure.CreateOneNeure(s.KeyPrefix, &neure.Neure{
			NeureType:              "normal",
			LastTimeActivate:       time.Now(),
			LastTimeResetNowWeight: time.Now(),
		})
		s.Neures = append(s.Neures, neureObj.ThisNeureId)
	}
}

func (s *Skin) struct2Byte() []byte {
	dataByte, err := json.Marshal(s)
	if err != nil {
		panic("json marshal error: " + err.Error())
	}
	return dataByte
}
