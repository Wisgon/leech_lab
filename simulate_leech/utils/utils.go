package utils

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"sync"
)

func SignalPass(entranceNeure *neure.Neure) {
	// todo:
}

func LoadFromMapByKeyPrefix[T body.Skin | body.Muscle | brain.Sense](dataMap *sync.Map, keyPrefix string, value *T) {
	dataByte := database.GetDataById(keyPrefix + config.PrefixNumSplitSymbol + "collection")
	Byte2Struct(dataByte, value)
	dataMap.Store(keyPrefix+config.PrefixNumSplitSymbol+"collection", value)
}

func StoreToMap[T body.Skin | body.Muscle | brain.Sense](dataMap *sync.Map, key string, value *T) {
	v, ok := dataMap.Load(key)
	var datas []*T
	if ok {
		datas = v.([]*T)
	}
	datas = append(datas, value)
	dataMap.Store(key, datas)
}

func Struct2Byte[T body.Skin | body.Muscle | brain.Sense](data *T) []byte {
	dataByte, err := json.Marshal(data)
	if err != nil {
		panic("json marshal error: " + err.Error())
	}
	return dataByte
}

func Byte2Struct[T body.Skin | body.Muscle | brain.Sense](neureByte []byte, data *T) {
	err := json.Unmarshal(neureByte, data)
	if err != nil {
		panic("json unmarshal error: " + err.Error())
	}
}
