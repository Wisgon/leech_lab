package neure

import (
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"log"
	"strings"
	"sync"
	"time"
)

var NeureMap = &sync.Map{}

func CheckNeureMap(stopSignal chan bool) {
	// check neure map, if length of neureMap bigger than MaxNeureMapNum, save it to db and remove it
	var breakSignal = false
	for {
		NeureMap.Range(func(key, value any) bool {
			keyString := key.(string)
			neureObj := value.(*Neure)
			lastTimeActivate := neureObj.LastTimeActivate
			if time.Since(lastTimeActivate) > config.InSyncNeureMapDuration {
				neureObj.UpdateNeure2DB()
				NeureMap.Delete(keyString)
			}
			select {
			case <-stopSignal:
				breakSignal = true
				return false // break neure map range loop
			default:
				return true
			}
		})
		if breakSignal {
			break
		}
		time.Sleep(config.InSyncNeureMapDuration + 1*time.Minute)
	}
}

func CreateOneNeure(keyPrefix string, neure *Neure) *Neure {
	uniqueNum := database.GetSeqNum(keyPrefix)
	key := keyPrefix + config.PrefixNumSplitSymbol + fmt.Sprint(uniqueNum)
	neure.ThisNeureId = key
	neure.SaveNeure2Db()
	NeureMap.Store(key, neure)
	return neure
}

func GetNeureById(id string) *Neure {
	np, ok := NeureMap.Load(id)
	if !ok {
		neure := &Neure{}
		neureByte := database.GetDataById(id)
		neure.Byte2Struct(neureByte)
		// store neure pointer to map
		NeureMap.Store(id, neure)
		return neure
	} else {
		neure := np.(*Neure)
		return neure
	}
}

func DeleteNeure(neure *Neure) {
	// delete the dendrites of next neures
	for _, synapse := range neure.Synapses {
		nextNeure := GetNeureById(synapse.NextNeureID)
		delete(nextNeure.NowLinkedDendritesIds, neure.ThisNeureId)
	}

	// delete the synapse of pre neures
	for dendriteId := range neure.NowLinkedDendritesIds {
		preNeure := GetNeureById(dendriteId)
		delete(preNeure.Synapses, neure.ThisNeureId)
	}

	database.DeleteNeure(neure.ThisNeureId)
	_, ok := NeureMap.Load(neure.ThisNeureId)
	if ok {
		NeureMap.Delete(neure.ThisNeureId)
	}
}

func TurnNeureBytes2Neures(neureBytes map[string][]byte) map[string]*Neure {
	neures := make(map[string]*Neure)
	for k, v := range neureBytes {
		neures[k] = &Neure{}
		neures[k].Byte2Struct(v)
	}
	return neures
}

func GetOtherTypeOfNeurePrefix(prefix string, neureType string) string {
	oldType := ""
	if strings.Contains(prefix, config.PrefixNeureType["common"]) {
		oldType = config.PrefixNeureType["common"]
	} else if strings.Contains(prefix, "regulate") {
		oldType = config.PrefixNeureType["regulate"]
	} else if strings.Contains(prefix, config.PrefixNeureType["inhibitory"]) {
		oldType = config.PrefixNeureType["inhibitory"]
	} else {
		log.Panic("unknow prefix neure type: " + prefix)
	}
	return strings.Replace(prefix, oldType, neureType, -1)
}
