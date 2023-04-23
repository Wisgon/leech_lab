package neure

import (
	"graph_robot/config"
	"sync"
	"time"
)

var NeureMap sync.Map

func CheckNeureMap() {
	// check neure map, if length of neureMap bigger than MaxNeureMapNum, save it to db and remove it
	for {
		NeureMap.Range(func(key, value any) bool {
			keyString := key.(string)
			neureObj := value.(*Neure)
			lastTimeActivate := neureObj.GetLastTimeActivate()
			if time.Since(lastTimeActivate) > config.InSyncNeureMapDuration {
				neureObj.UpdateNeure2DB()
				NeureMap.Delete(keyString)
			}
			return true
		})
		time.Sleep(config.InSyncNeureMapDuration + 1*time.Minute)
	}
}

func CreateOneNeure(keyPrefix string) (neure *Neure) {
	neure = &Neure{}
	neure.CreateNeureInDB(keyPrefix)
	return
}

func TurnNeureBytes2Neures(neureBytes *map[string][]byte) *map[string]*Neure {
	neures := make(map[string]*Neure)
	for k, v := range *neureBytes {
		neures[k] = &Neure{}
		neures[k].Byte2Struct(v)
	}
	return &neures
}

func RemoveUniqueValueFromSynapse(value string, s []*Synapse) []*Synapse {
	for i, v := range s {
		if v.NextNeureID == value {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}
