package brain

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"log"
	"strings"
	"sync"
)

type Sense struct {
	mu             sync.Mutex
	SenseNeureType string   `json:"a"`
	Position       string   `json:"c"`
	Neures         []string `json:"d"`
	KeyPrefix      string   `json:"e"`
}

func (s *Sense) GetNeures() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Neures
}

func (s *Sense) InitSense(wg *sync.WaitGroup) {
	defer wg.Done()
	s.createNeures()

	dataByte := s.struct2Byte()
	database.CreateData(dataByte, s.KeyPrefix+config.PrefixNumSplitSymbol+"collection")
}

func (s *Sense) createNeures() {
	var neureNum int
	if strings.Contains(s.SenseNeureType, "normal") {
		// sense neure is 10 times less than skin neure
		neureNum = config.EachSkinPositionSurfaceNeureNum / 10
	} else if strings.Contains(s.SenseNeureType, "bigger") {
		neureNum = config.EachSkinPositionDeeperNeureNum / 10
	} else if strings.Contains(s.SenseNeureType, "extremely") {
		neureNum = config.EachSkinPositionDeepestNeureNum / 10
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

func (s *Sense) struct2Byte() []byte {
	dataByte, err := json.Marshal(s)
	if err != nil {
		log.Panic("json marshal error: " + err.Error())
	}
	return dataByte
}

func IterSense(f func(senseNeureType, position string)) {
	for senseNeureType := range config.PrefixSkinAndSenseType {
		for position := range config.SkinAndSenseNeurePosition {
			f(senseNeureType, position)
		}
	}
}
