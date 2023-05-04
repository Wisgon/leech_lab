package brain

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"log"
	"sync"
	"time"
)

type Valuate struct {
	mu           sync.Mutex
	Source       string   `json:"a"`
	ValuateLevel string   `json:"b"`
	Neures       []string `json:"c"`
	KeyPrefix    string   `json:"d"`
}

func (v *Valuate) GetNeures() []string {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.Neures
}

func (v *Valuate) InitValuate(wg *sync.WaitGroup) {
	defer wg.Done()
	v.createNeures()

	dataByte := v.struct2Byte()
	database.CreateData(dataByte, v.KeyPrefix+config.PrefixNumSplitSymbol+"collection")
}

func (v *Valuate) struct2Byte() []byte {
	dataByte, err := json.Marshal(v)
	if err != nil {
		log.Panic("json marshal error: " + err.Error())
	}
	return dataByte
}

func (v *Valuate) createNeures() {
	for i := 0; i < config.EachValuateTypeNeureNum; i++ {
		neureObj := neure.CreateOneNeure(v.KeyPrefix, &neure.Neure{
			Synapses:               make(map[string]*neure.Synapse),
			NowLinkedDendritesIds:  make(map[string]struct{}),
			NeureType:              config.PrefixNeureType["common"],
			LastTimeActivate:       time.Now(),
			LastTimeResetNowWeight: time.Now(),
		})
		v.Neures = append(v.Neures, neureObj.ThisNeureId)
	}
}

func IterValuate(f func(valuateSource, valuateLevel string)) {
	for valuateSource := range config.PrefixValuateSource {
		for valuateLevel := range config.PrefixValuateLevel {
			f(valuateSource, valuateLevel)
		}
	}
}
