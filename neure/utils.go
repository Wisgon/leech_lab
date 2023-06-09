package neure

import (
	"context"
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"log"
	"strings"
	"sync"
	"time"
)

var NeureMap = &sync.Map{}
var SignalPassRecorder = make(chan map[string]interface{}, 1000) // for now the buffer is 1000

func CheckNeureMap(ctx context.Context) {
	// check neure map, if length of neureMap bigger than MaxNeureMapNum, save it to db and remove it
	var breakSignal = false
	go func() {
		<-ctx.Done()
		breakSignal = true
	}()
	for {
		if breakSignal {
			break
		}
		time.Sleep(config.InSyncNeureMapDuration + 1*time.Minute)
		log.Println("checking all neure~~~")
		NeureMap.Range(func(key, value any) bool {
			keyString := key.(string)
			neureObj := value.(*Neure)
			lastTimeActivate := neureObj.LastTimeActivate
			if time.Since(lastTimeActivate) > config.InSyncNeureMapDuration {
				neureObj.UpdateNeure2DB()
				neureObj.NeureSleep()
				NeureMap.Delete(keyString)
			}
			if breakSignal {
				return false
			} else {
				return true
			}
		})
	}
}

func CreateOneNeure(keyPrefix string, neure *Neure) *Neure {
	neure.Synapses = make(map[string]*Synapse)
	neure.NowLinkedDendritesIds = make(map[string]struct{})
	// don't set lastTimeActivate at now time.
	// now := time.Now()
	// neure.LastTimeActivate = now
	// neure.LastSignalTime = now
	bufferSize := config.SignalChannelBufferSizeDefault
	switch {
	case strings.Contains(keyPrefix, config.PrefixArea["sense"]) && strings.Contains(keyPrefix, "normal"):
		bufferSize = config.EachSkinPositionSurfaceNeureNum
	case strings.Contains(keyPrefix, config.PrefixArea["sense"]) && strings.Contains(keyPrefix, "bigger"):
		bufferSize = config.EachSkinPositionDeeperNeureNum
	case strings.Contains(keyPrefix, config.PrefixArea["sense"]) && strings.Contains(keyPrefix, "extremely"):
		bufferSize = config.EachSkinPositionDeepestNeureNum
	}
	neure.ChannelBufferSize = int32(bufferSize)
	neure.SignalChannel = make(chan map[string]interface{}, bufferSize)

	uniqueNum := database.GetSeqNum(keyPrefix)
	key := keyPrefix + config.PrefixNumSplitSymbol + fmt.Sprint(uniqueNum)
	neure.ThisNeureId = key
	neure.SaveNeure2Db()
	NeureMap.Store(key, neure)
	neure.WakeUpNeure() // sotre to map and then wakeup
	return neure
}

func GetNeureById(id string) *Neure {
	_, ok := NeureMap.Load(id)
	if !ok {
		neure := &Neure{}
		neureByte := database.GetDataById(id)
		neure.Byte2Struct(neureByte)
		// store neure pointer to map
		_, ok = NeureMap.Load(id) // check again because you must think about concurrent, this can make sure every goroutine load the same neure.
		if !ok {
			NeureMap.Store(id, neure)
			neure.WakeUpNeure() // sotre to map and then wakeup
		}
	}
	np, _ := NeureMap.Load(id)
	neure := np.(*Neure)
	return neure
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
