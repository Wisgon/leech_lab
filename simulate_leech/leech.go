package leech

import (
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/neure"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"graph_robot/simulate_leech/utils"
	"log"
	"strings"
	"sync"
)

type LeechBody struct {
	Organ *sync.Map
}

func (lb *LeechBody) InitBody(wg *sync.WaitGroup) {
	wg.Done()
	// init skin
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		body.IterSkin(func(skinNeureType, position string) {
			// only create common neure
			keyPrefix := config.PrefixArea["skin"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
			skin := body.Skin{
				SkinNeureType: skinNeureType,
				Position:      position,
				KeyPrefix:     keyPrefix,
			}
			wg.Add(1)
			go skin.InitSkin(wg)
			utils.StoreToMap(lb.Organ, keyPrefix+config.PrefixNumSplitSymbol+"collection", &skin)
		})
	}(wg)

	// init muscle
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		body.IterMuscle(func(movement string) {
			keyPrefix := config.PrefixArea["muscle"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + movement
			muscle := body.Muscle{
				MoveDirection: movement,
				KeyPrefix:     keyPrefix,
			}
			wg.Add(1)
			go muscle.InitMuscle(wg)
			utils.StoreToMap(lb.Organ, keyPrefix+config.PrefixNumSplitSymbol+"collection", &muscle)
		})
	}(wg)
}

func (lb *LeechBody) LoadBody(wg *sync.WaitGroup) {
	defer wg.Done()
	// load skin from database
	body.IterSkin(func(skinNeureType, position string) {
		keyPrefix := config.PrefixArea["skin"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
		skin := utils.LoadFromMapByKeyPrefix(lb.Organ, keyPrefix, &body.Skin{})
		// store more key because it's good to make more category for neures
		utils.StoreToMap(lb.Organ, config.PrefixArea["skin"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+skinNeureType, skin)
		utils.StoreToMap(lb.Organ, config.PrefixArea["skin"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+position, skin)
	})

	// load muscle from database
	body.IterMuscle(func(movement string) {
		keyPrefix := config.PrefixArea["muscle"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + movement
		muscle := utils.LoadFromMapByKeyPrefix(lb.Organ, keyPrefix, &body.Muscle{})
		utils.StoreToMap(lb.Organ, config.PrefixArea["muscle"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+movement, muscle)
	})

}

func (lb *LeechBody) Action(command string) {

}

func (lb *LeechBody) SenseFromEnv() { // get environment info
}

type LeechBrain struct {
	Area *sync.Map
}

func (lb *LeechBrain) InitBrain(wg *sync.WaitGroup) {
	defer wg.Done()
	// init sense
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		brain.IterSense(func(senseNeureType, position string) {
			keyPrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + position
			sense := brain.Sense{
				SenseNeureType: senseNeureType,
				Position:       position,
				KeyPrefix:      keyPrefix,
			}
			wg.Add(1)
			go sense.InitSense(wg)
			utils.StoreToMap(lb.Area, keyPrefix+config.PrefixNumSplitSymbol+"collection", &sense)
		})

		brain.IterValuate(func(valuateSource, valuateLevel string) {
			keyPrefix := config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + valuateSource + config.PrefixNameSplitSymbol + valuateLevel
			valuate := brain.Valuate{
				Source:       valuateSource,
				ValuateLevel: valuateLevel,
				KeyPrefix:    keyPrefix,
			}
			wg.Add(1)
			go valuate.InitValuate(wg)
			utils.StoreToMap(lb.Area, keyPrefix+config.PrefixNumSplitSymbol+"collection", &valuate)
		})
	}(wg)
}

func (lb *LeechBrain) LoadBrain(wg *sync.WaitGroup) {
	defer wg.Done()
	brain.IterSense(func(senseNeureType, position string) {
		keyPrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + position
		// store with some types so that can get one type very fast
		sense := utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &brain.Sense{})
		utils.StoreToMap(lb.Area, config.PrefixArea["sense"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+senseNeureType, sense)
		utils.StoreToMap(lb.Area, config.PrefixArea["sense"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+position, sense)
	})

	brain.IterValuate(func(valuateSource, valuateLevel string) {
		keyPrefix := config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + valuateSource + config.PrefixNameSplitSymbol + valuateLevel
		valuate := utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &brain.Valuate{})
		utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+valuateSource, valuate)
		utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+valuateLevel, valuate)

		// There are also regulate and inhibitory neure in valuate
		keyPrefix = config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["regulate"] + config.PrefixNameSplitSymbol + valuateSource + config.PrefixNameSplitSymbol + valuateLevel
		valuate = utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &brain.Valuate{})
		if valuate.KeyPrefix != "" {
			// means that key had been found
			utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["regulate"]+config.PrefixNameSplitSymbol+valuateSource, valuate)
			utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["regulate"]+config.PrefixNameSplitSymbol+valuateLevel, valuate)
		}

		keyPrefix = config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["inhibitory"] + config.PrefixNameSplitSymbol + valuateSource + config.PrefixNameSplitSymbol + valuateLevel
		valuate = utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &brain.Valuate{})
		if valuate.KeyPrefix != "" {
			// means that key had been found
			utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["inhibitory"]+config.PrefixNameSplitSymbol+valuateSource, valuate)
			utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["inhibitory"]+config.PrefixNameSplitSymbol+valuateLevel, valuate)
		}

	})
}

func (lb *LeechBrain) Sense2Action() (bodyAction string) {
	return
}

type Leech struct {
	Brain       *LeechBrain
	Body        *LeechBody
	EnvResponse chan map[string]interface{}
	EnvRequest  chan map[string]interface{}
}

func (l *Leech) InitLeech() {
	var wg sync.WaitGroup
	// init a leech
	wg.Add(2)
	go l.Brain.InitBrain(&wg)
	go l.Body.InitBody(&wg)

	wg.Wait()

	// after create neures, we can link neures, these neures is inborn neures, won't be deleted
	// first, link common type of skin and sense
	body.IterSkin(func(skinNeureType, position string) {
		skinPrefix := config.PrefixArea["skin"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
		sensePrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
		utils.LinkNeureGroups(
			utils.GetNeureIdsByKeyPrefix(l.Body.Organ, skinPrefix, &body.Skin{}),
			utils.GetNeureIdsByKeyPrefix(l.Brain.Area, sensePrefix, &brain.Sense{}),
			10, 1, config.PrefixNeureType["common"],
			func(synapseIds []string) (targetSynapseIds []string) { return },
		)
	})

	// senond, link common type of sense and muscle and valuate
	brain.IterSense(func(senseNeureType, position string) {
		opposite := utils.GetOpposite(position)
		// link to muscle
		sensePrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + position
		// sense link the opposite of muscle
		musclePrefix := config.PrefixArea["muscle"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + "move" + opposite
		utils.LinkNeureGroups(
			utils.GetNeureIdsByKeyPrefix(l.Body.Organ, sensePrefix, &brain.Sense{}),
			utils.GetNeureIdsByKeyPrefix(l.Brain.Area, musclePrefix, &body.Muscle{}),
			50, 1, config.PrefixNeureType["common"], // todo:sense和muscle的连接强度初始值50是否合理
			func(synapseIds []string) (targetSynapseIds []string) { return },
		)

		// link to valuate
		var valuatePrefix string
		switch {
		case strings.Contains(senseNeureType, "normal"):
			// normal sense is valuate0, not good not bad
			valuatePrefix = config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + config.PrefixValuateSource["sense"] + config.PrefixNameSplitSymbol + config.PrefixValuateLevel["valuate0"]
		case strings.Contains(senseNeureType, "bigger"):
			// bigger sense is valuate-1, a little uncomfortable
			valuatePrefix = config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + config.PrefixValuateSource["sense"] + config.PrefixNameSplitSymbol + config.PrefixValuateLevel["valuate-1"]
		case strings.Contains(senseNeureType, "extremely"):
			// extremely sense is valuate-2, painful
			valuatePrefix = config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + config.PrefixValuateSource["sense"] + config.PrefixNameSplitSymbol + config.PrefixValuateLevel["valuate-2"]
		default:
			log.Panic("unknow senseNeureType")
		}
		utils.LinkNeureGroups(
			utils.GetNeureIdsByKeyPrefix(l.Body.Organ, sensePrefix, &brain.Sense{}),
			utils.GetNeureIdsByKeyPrefix(l.Brain.Area, valuatePrefix, &brain.Valuate{}),
			101, 1, config.PrefixNeureType["common"],
			func(synapseIds []string) (targetSynapseIds []string) { return },
		)
	})

	// third, link valuate regulate to synapse that link sense and muscle
	groups := make(map[string][]string)
	brain.IterValuate(func(valuateSource, valuateLevel string) {
		valuateGroupKeyPrefix := config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + valuateSource + config.PrefixNameSplitSymbol + valuateLevel
		valuateGroupNeureIds := utils.GetNeureIdsByGroupName[*brain.Valuate](l.Brain.Area, valuateGroupKeyPrefix+config.PrefixNumSplitSymbol+"collection")
		switch valuateSource {
		case config.PrefixValuateSource["sense"]:
			if valuateLevel == config.PrefixValuateLevel["valuate1"] || valuateLevel == config.PrefixValuateLevel["valuate2"] {
				// sense has no valuate1 or valuate2 situlation
				return
			}
			// regulate and inhibitory neure must link all the synapse between sense and muscle
			newNeures := []string{}
			var groupName string
			brain.IterSense(func(senseNeureType, position string) {
				linkValuate2SenseFunction := func(neureType string) {
					senseKeyPrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + position + config.PrefixNumSplitSymbol + "collection"
					senseGroupNeureIds := utils.GetNeureIdsByGroupName[*brain.Sense](l.Brain.Area, senseKeyPrefix)
					opposite := utils.GetOpposite(position)
					musclePrefix := config.PrefixArea["muscle"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + "move" + opposite
					newNeureIds := utils.LinkNeureGroups(valuateGroupNeureIds, senseGroupNeureIds, 101, 1, neureType, func(synapseIds []string) (targetSynapseIds []string) {
						for _, synapseId := range synapseIds {
							if strings.Contains(synapseId, musclePrefix) {
								targetSynapseIds = append(targetSynapseIds, synapseId)
							}
						}
						return
					})
					newNeures = append(newNeures, newNeureIds...)
				}
				switch valuateLevel {
				case config.PrefixValuateLevel["valuate0"], config.PrefixValuateLevel["valuate-1"]:
					// valuate0 and -1 link inhibitory neure
					linkValuate2SenseFunction(config.PrefixNeureType["inhibitory"])
					groupName = neure.GetOtherTypeOfNeurePrefix(valuateGroupKeyPrefix, config.PrefixNeureType["inhibitory"])
					groups[groupName] = append(groups[groupName], newNeures...)
				case config.PrefixValuateLevel["valuate-2"]:
					// valuate-2 link regulate neure
					linkValuate2SenseFunction(config.PrefixNeureType["regulate"])
					groupName = neure.GetOtherTypeOfNeurePrefix(valuateGroupKeyPrefix, config.PrefixNeureType["regulate"])
					groups[groupName] = append(groups[groupName], newNeures...)
				}
			})
		default:
			log.Panic("wrong PrefixValuateSource:", valuateSource)
		}
	})
	// save all new neure that created during linking with regulate or inhibitory to database
	for groupName := range groups {
		splits := strings.Split(groupName, config.PrefixNameSplitSymbol)
		valuateLevel := splits[3] // make sure valueLevel is at index 3 of groupName!
		valueSource := splits[2]  // make sure valueSource is at index 2 of groupName!
		// store new neure to map
		valuate := &brain.Valuate{
			Source:       valueSource,
			ValuateLevel: valuateLevel,
			Neures:       groups[groupName],
			KeyPrefix:    groupName,
		}
		utils.StoreToMap(l.Brain.Area, valuate.KeyPrefix+config.PrefixNumSplitSymbol+"collection", valuate)

		dataByte := valuate.Struct2Byte()
		database.CreateData(dataByte, valuate.KeyPrefix+config.PrefixNumSplitSymbol+"collection")
	}

	// finally update all neures so that we can save the connect message into neure
	neure.NeureMap.Range(func(key, value any) bool {
		neureObj := value.(*neure.Neure)
		neureObj.UpdateNeure2DB()
		return true
	})
}

func (l *Leech) LoadLeech() {
	var wg sync.WaitGroup
	wg.Add(2)
	go l.Body.LoadBody(&wg)
	go l.Brain.LoadBrain(&wg)
	wg.Wait()
}

func (l *Leech) WakeUp() {
	maps := make(map[string]*sync.Map)
	maps["area"] = l.Brain.Area
	maps["organ"] = l.Body.Organ
	for {
		envResponse := <-l.EnvResponse
		event := envResponse["event"].(string)
		switch event {
		case "error":
			message := envResponse["message"].(string)
			log.Println("websocket error event: ", message)
			log.Panic("env response error")
		case "request_all_data":
			data := utils.AssembleMapDataToFront(maps)
			l.EnvRequest <- data
		case "request_part_data":
			parts := envResponse["message"].(map[string]interface{})
			log.Println("here comes a prefix: ", parts)
			data := utils.AssemblePartOfMapDataToFront(maps, parts)
			l.EnvRequest <- data
		case "link":
			linkCondition := envResponse["message"].(map[string]interface{})
			// linkCondition is a map {source:"xxx", strength:10, target:"yyy", link_type: "common", synapse_id:""}
			log.Println("websocket link event: ", linkCondition)
			link_type := linkCondition["link_type"].(string)
			if link_type == config.PrefixNeureType["common"] {
				utils.LinkTwoNeures(linkCondition)
			} else {
				// todo: decide which dataMap to pass, for now just l.Brain.Area
				utils.LinkTwoNeures(linkCondition)
			}
			data := make(map[string]interface{})
			l.EnvRequest <- data
		case "env_info":
			linkCondition := envResponse["message"].(map[string]interface{})
			log.Printf("get evn info: %+v\n", linkCondition)
			// todo: do something with env info
		default:
			log.Println("unknow event:", event)
		}
	}
}
