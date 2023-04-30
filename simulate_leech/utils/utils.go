package utils

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CreatureParts interface {
	*body.Skin | *body.Muscle | *brain.Sense
	GetNeures() []string
}

func SignalPass(entranceNeure *neure.Neure) {
	// todo:
}

func LinkTwoNeures(linkCondition map[string]interface{}) {
	source, target := linkCondition["source"].(string), linkCondition["target"].(string)
	linkType := linkCondition["link_type"].(string)
	var strength float64
	var err error
	switch strengthType := linkCondition["strength"].(type) {
	case string:
		strength, err = strconv.ParseFloat(strengthType, 64)
		if err != nil {
			log.Println("error: parse strength fail, link fail")
			return
		}
	case float32:
		strength = float64(strengthType)
	}
	if linkType == "common" {
		neureSource := neure.GetNeureById(source)
		neureSource.ConnectNextNuere(&neure.Synapse{
			NextNeureID:  target,
			LinkStrength: float32(strength),
			SynapseNum:   1,
		})
	} else {
		if linkType != "regulate" && linkType != "inhibitory" {
			log.Panic("wrong neure type:" + linkType)
		}
		// need to create a new neure
		newNeurePrefix := neure.GetOtherTypeOfNeurePrefix(source, linkType)
		regulateNeure := neure.CreateOneNeure(newNeurePrefix, &neure.Neure{
			NeureType:              linkType,
			LastTimeActivate:       time.Now(),
			LastTimeResetNowWeight: time.Now(),
		})
		neureSource := neure.GetNeureById(source)
		neureSource.ConnectNextNuere(&neure.Synapse{
			NextNeureID:  regulateNeure.ThisNeureId,
			LinkStrength: float32(strength),
			SynapseNum:   1,
		})
	}

}

func LinkNeureGroups(sourceNeures []string, targetNeures []string, strength float32, synapseNum int32, linkType string) {
	for _, neureId := range sourceNeures {
		linkCondition := make(map[string]interface{})
		nextNeureId := targetNeures[rand.Intn(len(targetNeures))] // link to random neure in targetNeures
		linkCondition["source"] = neureId
		linkCondition["target"] = nextNeureId
		linkCondition["strength"] = strength
		linkCondition["link_type"] = linkType
		LinkTwoNeures(linkCondition)
	}
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
			case "common":
				neureType = "c"
			case "regulate":
				neureType = "r"
			case "inhibitory":
				neureType = "i"
			default:
				log.Panic("wrong neure type:" + neureObj.NeureType)
			}
			link["neure_type"] = neureType
			(*links) = append((*links), link)
		}
	}
}

func GetNeureIdsByKeyPrefix[T CreatureParts](dataMap *sync.Map, keyPrefix string, value T) []string {
	LoadFromMapByKeyPrefix(dataMap, keyPrefix, value)
	return value.GetNeures()
}

func LoadFromMapByKeyPrefix[T CreatureParts](dataMap *sync.Map, keyPrefix string, value T) {
	dataByte := database.GetDataById(keyPrefix + config.PrefixNumSplitSymbol + "collection")
	Byte2Struct(dataByte, value)
	dataMap.Store(keyPrefix+config.PrefixNumSplitSymbol+"collection", value)
}

func StoreToMap[T CreatureParts](dataMap *sync.Map, key string, value T) {
	v, ok := dataMap.Load(key)
	var datas []T
	if ok {
		datas = v.([]T)
	}
	datas = append(datas, value)
	dataMap.Store(key, datas)
}

func Struct2Byte[T CreatureParts](data T) []byte {
	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Panic("json marshal error: " + err.Error())
	}
	return dataByte
}

func Byte2Struct[T CreatureParts](neureByte []byte, data T) {
	err := json.Unmarshal(neureByte, data)
	if err != nil {
		log.Panic("json unmarshal error: " + err.Error())
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

func GetOpposite(position string) (opposite string) {
	opposite = strings.Replace(position, "left", "Right", -1)
	opposite = strings.Replace(opposite, "Front", "Back", -1)
	opposite = strings.Replace(opposite, "Up", "Down", -1)
	opposite = strings.Replace(opposite, "right", "Left", -1)
	opposite = strings.Replace(opposite, "Back", "Front", -1)
	opposite = strings.Replace(opposite, "Down", "Up", -1)
	return
}
