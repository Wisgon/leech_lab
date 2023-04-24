package neure

import (
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/utils"
	"sync"
	"time"
)

var NeureMap sync.Map // sync.Map必须储存值，而不能储存引用，而且每次有更新都要用 NeureMap.Store() 方法来更新

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
		neureByte := database.GetNeure(id)
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
		utils.RemoveUniqueValueFromSlice(neure.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
	}

	// delete the synapse of pre neures
	for _, dendriteId := range neure.NowLinkedDendritesIds {
		preNeure := GetNeureById(dendriteId)
		preNeure.RemoveSynapseByNextId(neure.ThisNeureId)
	}

	database.DeleteNeure(neure.ThisNeureId)
	_, ok := NeureMap.Load(neure.ThisNeureId)
	if ok {
		NeureMap.Delete(neure.ThisNeureId)
	}
}

func TurnNeureBytes2Neures(neureBytes *map[string][]byte) *map[string]*Neure {
	neures := make(map[string]*Neure)
	for k, v := range *neureBytes {
		neures[k] = &Neure{}
		neures[k].Byte2Struct(v)
	}
	return &neures
}
