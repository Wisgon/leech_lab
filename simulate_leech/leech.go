package leech

import (
	"graph_robot/config"
	"graph_robot/neure"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"graph_robot/simulate_leech/utils"
	"log"
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
		skin := body.Skin{
			SkinNeureType: skinNeureType,
			Position:      position,
			KeyPrefix:     keyPrefix,
		}
		utils.LoadFromMapByKeyPrefix(lb.Organ, keyPrefix, &skin)
		// store more key because it's good to make more category for neures
		utils.StoreToMap(lb.Organ, config.PrefixArea["skin"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+skinNeureType, &skin)
		utils.StoreToMap(lb.Organ, config.PrefixArea["skin"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+position, &skin)
	})

	// load muscle from database
	body.IterMuscle(func(movement string) {
		keyPrefix := config.PrefixArea["muscle"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + movement
		muscle := body.Muscle{
			MoveDirection: movement,
			KeyPrefix:     keyPrefix,
		}
		utils.LoadFromMapByKeyPrefix(lb.Organ, keyPrefix, &muscle)
		utils.StoreToMap(lb.Organ, config.PrefixArea["muscle"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+movement, &muscle)
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
		sense := brain.Sense{
			SenseNeureType: senseNeureType,
			Position:       position,
			KeyPrefix:      keyPrefix,
		}
		// store with some types so that can get one type very fast
		utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &sense)
		utils.StoreToMap(lb.Area, config.PrefixArea["sense"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+senseNeureType, &sense)
		utils.StoreToMap(lb.Area, config.PrefixArea["sense"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+position, &sense)
	})

	brain.IterValuate(func(valuateSource, valuateLevel string) {
		keyPrefix := config.PrefixArea["valuate"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + valuateSource + config.PrefixNameSplitSymbol + valuateLevel
		valuate := brain.Valuate{
			Source:       valuateSource,
			ValuateLevel: valuateLevel,
			KeyPrefix:    keyPrefix,
		}
		utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &valuate)
		utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+valuateSource, &valuate)
		utils.StoreToMap(lb.Area, config.PrefixArea["valuate"]+config.PrefixNameSplitSymbol+config.PrefixNeureType["common"]+config.PrefixNameSplitSymbol+valuateLevel, &valuate)
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

	// link neures, these neures is inborn neures, won't be deleted
	// first, link common type of skin and sense
	body.IterSkin(func(skinNeureType, position string) {
		skinPrefix := config.PrefixArea["skin"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
		sensePrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
		utils.LinkNeureGroups(
			utils.GetNeureIdsByKeyPrefix(l.Body.Organ, skinPrefix, &body.Skin{}),
			utils.GetNeureIdsByKeyPrefix(l.Brain.Area, sensePrefix, &brain.Sense{}),
			10, 1, config.PrefixNeureType["common"],
		)
	})

	// senond, link common type of sense and muscle
	brain.IterSense(func(senseNeureType, position string) {
		opposite := utils.GetOpposite(position)
		log.Println("position:", position, " opposite:", opposite)
		sensePrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + position
		// sense link the opposite of muscle
		musclePrefix := config.PrefixArea["muscle"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + "move" + opposite
		utils.LinkNeureGroups(
			utils.GetNeureIdsByKeyPrefix(l.Body.Organ, sensePrefix, &brain.Sense{}),
			utils.GetNeureIdsByKeyPrefix(l.Brain.Area, musclePrefix, &body.Muscle{}),
			50, 1, config.PrefixNeureType["common"], // todo:sense和muscle的连接强度初始值50是否合理
		)
	})

	// third, link regulate type of painful type to the synapse of sense to muscle
	// todo: connect to valuate area，思考神经元激活频率，数量，weight之间的关系
	// brain.IterSense(func(senseNeureType, senseType, position string) {
	// 	keyPrefix := config.PrefixArea["sense"] + config.PrefixNameSplitSymbol + config.PrefixNeureType["common"] + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + senseType + config.PrefixNameSplitSymbol + position
	// 	senseCollection, ok := l.Brain.Area.Load(keyPrefix + config.PrefixNumSplitSymbol + "collection")
	// 	if !ok {
	// 		return
	// 	}
	// 	sense := senseCollection.(*brain.Sense)
	// 	for _, sense := range sense.Neures {
	// 		senseNeure := neure.GetNeureById(sense)
	// 		for _, senseSynapse := range senseNeure.Synapses {
	// 			linkCondition := make(map[string]interface{})
	// 			linkCondition["source"] = senseNeure.ThisNeureId
	// 			linkCondition["target"]
	// 		}

	// 	}
	// })

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
			// linkCondition is a map {source:"xxx", strength:10, target:"yyy", link_type: "n", synapse_id:""}
			log.Println("websocket link event: ", linkCondition)
			utils.LinkTwoNeures(linkCondition)
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
