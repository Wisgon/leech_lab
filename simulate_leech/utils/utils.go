package utils

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"log"
	"strconv"
	"strings"
	"sync"
)

func SignalPass(entranceNeure *neure.Neure) {
	// todo:
}

func LinkNeures(linkCondition map[string]interface{}) {
	source, target := linkCondition["source"].(string), linkCondition["target"].(string)
	strengthStr := linkCondition["strength"].(string)
	strength, err := strconv.ParseFloat(strengthStr, 64)
	if err != nil {
		log.Println("error: parse strength fail, link fail")
		return
	}
	neureSource := neure.GetNeureById(source)
	synapse := neure.Synapse{
		NextNeureID:  target,
		LinkStrength: float32(strength),
		SynapseNum:   1,
	}
	neureSource.ConnectNextNuere(&synapse)
}

func AssembleLinkData(keyStr string, neures []string, groups *map[string][]string, links *[]map[string]interface{}) {
	for _, v := range neures {
		(*groups)[keyStr] = append((*groups)[keyStr], v)
		neureObj := neure.GetNeureById(v)
		for _, s := range neureObj.Synapses {
			link := make(map[string]interface{})
			link["source"] = v
			link["target"] = s.NextNeureID
			link["link_strength"] = s.LinkStrength
			link["synapse_num"] = s.SynapseNum
			neureType := ""
			switch neureObj.NeureType {
			case "normal":
				neureType = "n"
			case "regulate":
				neureType = "r"
			case "inhibitory":
				neureType = "i"
			default:
				panic("wrong neure type:" + neureObj.NeureType)
			}
			link["neure_type"] = neureType
			(*links) = append((*links), link)
		}
	}
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

func AssembleMapDataToFront(area *sync.Map, organ *sync.Map) map[string]interface{} {
	data := make(map[string]interface{})
	links := []map[string]interface{}{}
	nodes := []map[string]interface{}{}
	groups := make(map[string][]string)
	area.Range(func(key, value any) bool {
		keyStr := key.(string)
		if strings.Contains(keyStr, "collection") {
			switch collection := value.(type) {
			case *body.Skin:
				AssembleLinkData(keyStr, collection.Neures, &groups, &links)
			case *body.Muscle:
				AssembleLinkData(keyStr, collection.Neures, &groups, &links)
			case *brain.Sense:
				AssembleLinkData(keyStr, collection.Neures, &groups, &links)
			}
		}
		return true
	})
	organ.Range(func(key, value any) bool {
		keyStr := key.(string)
		if strings.Contains(keyStr, "collection") {
			switch collection := value.(type) {
			case *body.Skin:
				AssembleLinkData(keyStr, collection.Neures, &groups, &links)
			case *body.Muscle:
				AssembleLinkData(keyStr, collection.Neures, &groups, &links)
			case *brain.Sense:
				AssembleLinkData(keyStr, collection.Neures, &groups, &links)
			}
		}
		return true
	})
	groupNum := 1
	for _, group := range groups {
		for _, neureId := range group {
			node := make(map[string]interface{})
			node["id"] = neureId
			node["group"] = groupNum
			nodes = append(nodes, node)
		}
		groupNum += 1

	}
	data["links"] = links
	data["nodes"] = nodes
	return data
}
