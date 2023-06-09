package body

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"log"
	"strings"
	"sync"
)

type Skin struct {
	mu            sync.Mutex
	Position      string   `json:"a"` // position of this part of Skin
	SkinNeureType string   `json:"b"` // type of skin neure
	Neures        []string `json:"c"` // neure id of this position
	KeyPrefix     string   `json:"d"`
}

func (s *Skin) GetNeures() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Neures
}

func (s *Skin) InitSkin(wg *sync.WaitGroup) {
	defer wg.Done()
	s.createNeures()

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
		log.Panic("Unknow skin neure type")
	}
	for i := 0; i < neureNum; i++ {
		neureObj := neure.CreateOneNeure(s.KeyPrefix, &neure.Neure{
			NeureType: config.PrefixNeureType["common"],
		})
		s.Neures = append(s.Neures, neureObj.ThisNeureId)
	}
}

func (s *Skin) struct2Byte() []byte {
	dataByte, err := json.Marshal(s)
	if err != nil {
		log.Panic("json marshal error: " + err.Error())
	}
	return dataByte
}

func IterSkin(f func(skinNeureType, position string)) {
	for skinNeureType := range config.PrefixSkinAndSenseType {
		for position := range config.SkinAndSenseNeurePosition {
			f(skinNeureType, position)
		}
	}
}
