package utils

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"graph_robot/utils"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

type CreatureParts interface {
	*body.Skin | *body.Muscle | *brain.Sense | *brain.Valuate
	GetNeures() []string
}

func LinkTwoNeures(linkCondition map[string]interface{}) (regulateNeure *neure.Neure) {
	source, target, synapse_id := linkCondition["source"].(string), linkCondition["target"].(string), linkCondition["synapse_id"].(string)
	linkType := linkCondition["link_type"].(string)
	var strength float32
	switch strengthUnknowType := linkCondition["strength"].(type) {
	case string:
		strength64, err := strconv.ParseFloat(strengthUnknowType, 64)
		if err != nil {
			log.Panic("error: parse strength fail, link fail:", err)
		}
		strength = float32(strength64)
	case float32:
		strength = strengthUnknowType
	case float64:
		strength = float32(strengthUnknowType)
	default:
		log.Panic("error strength type:", linkCondition["strength"])
	}
	var synapseNum int32
	switch synapseNumUnknowType := linkCondition["synapse_num"].(type) {
	case string:
		synapseNum64, err := strconv.ParseInt(synapseNumUnknowType, 0, 64)
		if err != nil {
			log.Panic("error: parse synapse num error: ", err)
		}
		synapseNum = int32(synapseNum64)
	case int:
		synapseNum = int32(synapseNumUnknowType)
	case int32:
		synapseNum = synapseNumUnknowType
	case int64:
		synapseNum = int32(synapseNumUnknowType)
	default:
		log.Panic("error synapse num type:", linkCondition["synapse_num"])
	}
	var hibituationbility bool
	switch hibituationbilityUnknowType := linkCondition["hibituationbility"].(type) {
	case string:
		if hibituationbilityUnknowType == "true" {
			hibituationbility = true
		} else if hibituationbilityUnknowType == "false" {
			hibituationbility = false
		} else {
			log.Panic("error hibituationbility value:", linkCondition["hibituationbility"])
		}
	case bool:
		hibituationbility = hibituationbilityUnknowType
	default:
		log.Panic("error hibituationbility type:", linkCondition["hibituationbility"])
	}
	if linkType == config.PrefixNeureType["common"] {
		neureSource := neure.GetNeureById(source)
		neureSource.ConnectNextNuere(&neure.Synapse{
			NextNeureID:       target,
			LinkStrength:      float32(strength),
			SynapseNum:        synapseNum,
			Hibituationbility: hibituationbility,
		})
	} else {
		if linkType != config.PrefixNeureType["regulate"] && linkType != config.PrefixNeureType["inhibitory"] {
			log.Panic("wrong neure type:" + linkType)
		}
		if synapse_id == "" {
			log.Panic("synapse_id must not be empty")
		}
		// need to create a new neure
		sourcePrefix := strings.Split(source, config.PrefixNumSplitSymbol)[0]
		newNeurePrefix := neure.GetOtherTypeOfNeurePrefix(sourcePrefix, linkType)
		regulateNeure = neure.CreateOneNeure(newNeurePrefix, &neure.Neure{
			NeureType: linkType,
		})
		neureSource := neure.GetNeureById(source)
		// first, connect source and regulate neure
		neureSource.ConnectNextNuere(&neure.Synapse{
			NextNeureID:       regulateNeure.ThisNeureId,
			LinkStrength:      float32(strength),
			SynapseNum:        synapseNum,
			Hibituationbility: hibituationbility,
		})
		// second, connect regulate neure to target synapse
		regulateNeure.ConnectNextNuere(&neure.Synapse{
			NextNeureID:        target,
			LinkStrength:       float32(strength),
			SynapseNum:         synapseNum,
			Hibituationbility:  hibituationbility,
			NextNeureSynapseId: synapse_id,
		})
	}
	return
}

func LinkNeureGroups(sourceNeures []string, targetNeures []string, linkCondition map[string]interface{}, fu func(synapseIds []string) (targetSynapseIds []string), linkRandomly bool) (newNeureIds []string) {
	if len(sourceNeures) == 0 || len(targetNeures) == 0 {
		return
	}
	linkType := linkCondition["link_type"].(string)
	everyTargetLinkNum := int(len(sourceNeures) / len(targetNeures))
	if everyTargetLinkNum == 0 {
		everyTargetLinkNum = 1
	}
	nowTargetIndex := -1
	// f is a function that give the synapseIds which get from target neures and return target synapseIds
	for i, neureId := range sourceNeures {
		var nextNeureId string
		if linkRandomly {
			nextNeureId = targetNeures[rand.Intn(len(targetNeures))] // link to random neure in targetNeures
		} else {
			if i%everyTargetLinkNum == 0 {
				nowTargetIndex += 1
			}
			nextNeureId = targetNeures[nowTargetIndex]
		}
		linkCondition["source"] = neureId
		linkCondition["target"] = nextNeureId
		if linkType == config.PrefixNeureType["common"] {
			linkCondition["synapse_id"] = ""
			LinkTwoNeures(linkCondition)
		} else {
			// inhibitory and regulate neure
			for _, nextNeureId := range targetNeures {
				linkCondition["target"] = nextNeureId
				// regulate and inhibitory must link all target neure
				nextNeure := neure.GetNeureById(nextNeureId)
				targetSynapseIds := fu(utils.GetMapKeys(nextNeure.Synapses))
				for _, targetSynapseId := range targetSynapseIds {
					linkCondition["synapse_id"] = targetSynapseId
					newNeure := LinkTwoNeures(linkCondition)
					newNeureIds = append(newNeureIds, newNeure.ThisNeureId)
				}
			}
		}
	}
	return
}

func assembleLinkData(neures []string, groups map[string][]string, links *[]map[string]interface{}, dendritesFlag bool) {
	for _, v := range neures {
		neureGroupName := strings.Split(v, config.PrefixNumSplitSymbol)[0]
		groups[neureGroupName] = append(groups[neureGroupName], v)
		neureObj := neure.GetNeureById(v)
		for _, s := range neureObj.Synapses {
			if dendritesFlag {
				synapseGroupName := strings.Split(s.NextNeureID, config.PrefixNumSplitSymbol)[0]
				groups[synapseGroupName] = append(groups[synapseGroupName], s.NextNeureID)
			}
			link := make(map[string]interface{})
			link["source"] = v
			link["target"] = s.NextNeureID
			link["link_strength"] = s.LinkStrength
			link["synapse_num"] = s.SynapseNum
			neureType := ""
			switch neureObj.NeureType {
			case config.PrefixNeureType["common"]:
				neureType = "c"
			case config.PrefixNeureType["regulate"]:
				neureType = "r"
			case config.PrefixNeureType["inhibitory"]:
				neureType = "i"
			default:
				log.Panic("wrong neure type:" + neureObj.NeureType)
			}
			link["neure_type"] = neureType
			(*links) = append((*links), link)
		}
		if dendritesFlag {
			for dendriteId := range neureObj.NowLinkedDendritesIds {
				dendriteGroupName := strings.Split(dendriteId, config.PrefixNumSplitSymbol)[0]
				groups[dendriteGroupName] = append(groups[dendriteGroupName], dendriteId)
				dendriteNeure := neure.GetNeureById(dendriteId)
				link := make(map[string]interface{})
				link["source"] = dendriteNeure.ThisNeureId
				link["target"] = neureObj.ThisNeureId
				var synapse *neure.Synapse
				for _, s := range dendriteNeure.Synapses {
					if s.NextNeureID == neureObj.ThisNeureId {
						synapse = s
					}
				}
				link["link_strength"] = synapse.LinkStrength
				link["synapse_num"] = synapse.SynapseNum
				neureType := ""
				switch dendriteNeure.NeureType {
				case config.PrefixNeureType["common"]:
					neureType = "c"
				case config.PrefixNeureType["regulate"]:
					neureType = "r"
				case config.PrefixNeureType["inhibitory"]:
					neureType = "i"
				default:
					log.Panic("wrong neure type:" + dendriteNeure.NeureType)
				}
				link["neure_type"] = neureType
				(*links) = append((*links), link)
			}
		}
	}
}

func GetNeureIdsByGroupName[T CreatureParts](dataMap *sync.Map, groupName string) []string {
	values := LoadFromMapByGroupName[T](dataMap, groupName)
	neureUnique := make(map[string]struct{})
	for _, value := range values {
		neures := value.GetNeures()
		for _, neure := range neures {
			neureUnique[neure] = struct{}{}
		}
	}
	return utils.GetMapKeys(neureUnique)
}

func LoadFromMapByGroupName[T CreatureParts](dataMap *sync.Map, groupName string) (values []T) {
	mapValue, ok := dataMap.Load(groupName)
	if ok {
		values = mapValue.([]T)
	} else {
		log.Panic("groupName:", groupName, " is not in data map")
	}
	return
}

func GetNeureIdsByKeyPrefix[T CreatureParts](dataMap *sync.Map, keyPrefix string, value T) []string {
	value = LoadFromMapByKeyPrefix(dataMap, keyPrefix, value)
	return value.GetNeures()
}

func LoadFromMapByKeyPrefix[T CreatureParts](dataMap *sync.Map, keyPrefix string, value T) T {
	mapKey := keyPrefix + config.PrefixNumSplitSymbol + "collection"
	// first try to load from map
	mapValue, ok := dataMap.Load(keyPrefix)
	var values []T // stay as same as the data struct saved in function StoreToMap
	if ok {
		values = mapValue.([]T)
		if len(values) == 0 {
			log.Panic("key:", keyPrefix, " is empty slice")
		}
		return values[0] // there is only one skin in one particular prefix
	} else {
		dataByte := database.GetDataById(mapKey)
		if len(dataByte) == 0 {
			// key not found
			return value
		}
		Byte2Struct(dataByte, value)
		values = append(values, value)
		dataMap.Store(keyPrefix+config.PrefixNumSplitSymbol+"collection", values)
		return value
	}
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

func checkIfInParts(partsStr []string, keyPrefix string) bool {
	for _, partStr := range partsStr {
		// strings.Contains contain "" will return true
		if !strings.Contains(keyPrefix, partStr) {
			return false
		}
	}
	return true
}

func getCollections(parts map[string]interface{}) (collections []string, partsStr []string) {
	prefix := parts["area"].(string) + config.PrefixNameSplitSymbol + parts["neure_type"].(string)
	if skin_sense_type, ok := parts["skin_sense_type"]; ok {
		value := skin_sense_type.(string)
		partsStr = append(partsStr, value)
		collections = append(collections, prefix+config.PrefixNameSplitSymbol+value)
	}
	if skin_sense_position, ok := parts["skin_sense_position"]; ok {
		value := skin_sense_position.(string)
		partsStr = append(partsStr, value)
		collections = append(collections, prefix+config.PrefixNameSplitSymbol+value)
	}
	if movements, ok := parts["movements"]; ok {
		value := movements.(string)
		partsStr = append(partsStr, value)
		collections = append(collections, prefix+config.PrefixNameSplitSymbol+value)
	}
	if valuate_source, ok := parts["valuate_source"]; ok {
		value := valuate_source.(string)
		partsStr = append(partsStr, value)
		collections = append(collections, prefix+config.PrefixNameSplitSymbol+value)
	}
	if valuate_level, ok := parts["valuate_level"]; ok {
		value := valuate_level.(string)
		partsStr = append(partsStr, value)
		collections = append(collections, prefix+config.PrefixNameSplitSymbol+value)
	}
	return
}

func removeRepeatFromCollections(partsStr []string, maps map[string]*sync.Map, collections []string) (uniqueNeures []string) {
	// different collection may have repeat element
	var neures []string
	neureUnique := make(map[string]struct{})
	for _, collection := range collections {
		area := strings.Split(collection, config.PrefixNameSplitSymbol)[0]
		switch area {
		case config.PrefixArea["skin"]:
			value, ok := maps["organ"].Load(collection)
			if ok {
				skins := value.([]*body.Skin)
				for _, skin := range skins {
					neures = append(neures, skin.Neures...)
				}
			}
		case config.PrefixArea["sense"]:
			value, ok := maps["area"].Load(collection)
			if ok {
				senses := value.([]*brain.Sense)
				for _, sense := range senses {
					neures = append(neures, sense.Neures...)
				}
			}
		case config.PrefixArea["muscle"]:
			value, ok := maps["organ"].Load(collection)
			if ok {
				muscles := value.([]*body.Muscle)
				for _, muscle := range muscles {
					neures = append(neures, muscle.Neures...)
				}
			}
		case config.PrefixArea["valuate"]:
			value, ok := maps["area"].Load(collection)
			if ok {
				valuates := value.([]*brain.Valuate)
				for _, valuate := range valuates {
					neures = append(neures, valuate.Neures...)
				}
			}
		}
	}
	for _, n := range neures {
		neureUnique[n] = struct{}{}
	}
	for key := range neureUnique {
		if !checkIfInParts(partsStr, key) {
			continue
		}
		uniqueNeures = append(uniqueNeures, key)
	}
	return
}

func AssemblePartOfMapDataToFront(maps map[string]*sync.Map, parts map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	links := []map[string]interface{}{}
	nodes := []map[string]interface{}{}
	groups := make(map[string][]string)

	collections, partsStr := getCollections(parts)
	uniqueNeures := removeRepeatFromCollections(partsStr, maps, collections)
	assembleLinkData(uniqueNeures, groups, &links, true)
	for groupName, group := range groups {
		// same neure won't appear in different group
		uniqueNodeId := make(map[string]struct{})
		for _, neureId := range group {
			_, ok := uniqueNodeId[neureId] // make sure nodeId unique
			if ok {
				// nodeId has been added
				continue
			}
			node := make(map[string]interface{})
			node["id"] = neureId
			node["group"] = groupName
			nodes = append(nodes, node)
			uniqueNodeId[neureId] = struct{}{}
		}
	}
	data["links"] = links
	data["nodes"] = nodes
	return data
}

func AssembleMapDataToFront(maps map[string]*sync.Map) map[string]interface{} {
	// get all data
	data := make(map[string]interface{})
	links := []map[string]interface{}{}
	nodes := []map[string]interface{}{}
	groups := make(map[string][]string)
	maps["area"].Range(func(key, value any) bool {
		keyStr := key.(string)
		if strings.Contains(keyStr, "collection") {
			switch collection := value.(type) {
			case *body.Skin:
				assembleLinkData(collection.Neures, groups, &links, false)
			case *body.Muscle:
				assembleLinkData(collection.Neures, groups, &links, false)
			case *brain.Sense:
				assembleLinkData(collection.Neures, groups, &links, false)
			}
		}
		return true
	})
	maps["organ"].Range(func(key, value any) bool {
		keyStr := key.(string)
		if strings.Contains(keyStr, "collection") {
			switch collection := value.(type) {
			case *body.Skin:
				assembleLinkData(collection.Neures, groups, &links, false)
			case *body.Muscle:
				assembleLinkData(collection.Neures, groups, &links, false)
			case *brain.Sense:
				assembleLinkData(collection.Neures, groups, &links, false)
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
	opposite = position
	if strings.Contains(position, "left") {
		opposite = strings.Replace(position, "left", "Right", -1)
	} else {
		opposite = strings.Replace(opposite, "right", "Left", -1)
	}
	if strings.Contains(position, "Front") {
		opposite = strings.Replace(opposite, "Front", "Back", -1)
	} else {
		opposite = strings.Replace(opposite, "Back", "Front", -1)
	}
	if strings.Contains(position, "Up") {
		opposite = strings.Replace(opposite, "Up", "Down", -1)
	} else {
		opposite = strings.Replace(opposite, "Down", "Up", -1)
	}
	return
}

func ParseResult(resultNeureIds []string) {
	// todo: parse result and get output, whatever it is.
}
