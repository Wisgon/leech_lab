package leech

import (
	"graph_robot/config"
	"graph_robot/interact"
	"graph_robot/simulate_leech/body"
	"graph_robot/simulate_leech/brain"
	"graph_robot/simulate_leech/utils"
	"strings"
	"sync"
)

type LeechBody struct {
	Organ *sync.Map
}

func (lb *LeechBody) InitBody(wg *sync.WaitGroup, processController *sync.Map) {
	// init skin
	wg.Add(1)
	go func(wg *sync.WaitGroup, processController *sync.Map) {
		defer wg.Done()

		for _, skinNeureType := range config.PrefixSkinAndSenseType {
			for _, position := range config.SkinAndSenseNeurePosition {
				keyPrefix := "skin" + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
				skin := body.Skin{
					SkinNeureType: skinNeureType,
					Position:      position,
					KeyPrefix:     keyPrefix,
				}
				wg.Add(1)
				go skin.InitSkin(wg, processController)
			}
		}
	}(wg, processController)

	// init muscle
	wg.Add(1)
	go func(wg *sync.WaitGroup, processController *sync.Map) {
		defer wg.Done()

		for _, movement := range config.Movements {
			keyPrefix := "muscle" + config.PrefixNameSplitSymbol + movement
			muscle := body.Muscle{
				MoveDirection: movement,
				KeyPrefix:     keyPrefix,
			}

			utils.StoreToMap(lb.Organ, movement, &muscle)
			utils.StoreToMap(lb.Organ, keyPrefix, &muscle)
			wg.Add(1)
			go muscle.InitMuscle(wg, processController)
		}
	}(wg, processController)
}

func (lb *LeechBody) LoadBody(wg *sync.WaitGroup) {
	defer wg.Done()
	// load skin from database
	for _, skinNeureType := range config.PrefixSkinAndSenseType {
		for _, position := range config.SkinAndSenseNeurePosition {
			keyPrefix := "skin" + config.PrefixNameSplitSymbol + skinNeureType + config.PrefixNameSplitSymbol + position
			skin := body.Skin{
				SkinNeureType: skinNeureType,
				Position:      position,
				KeyPrefix:     keyPrefix,
			}
			utils.LoadFromMapByKeyPrefix(lb.Organ, keyPrefix, &skin)
			utils.StoreToMap(lb.Organ, skinNeureType, &skin)
			utils.StoreToMap(lb.Organ, position, &skin)
		}
	}

	// load muscle from database
	for _, movement := range config.Movements {
		keyPrefix := "muscle" + config.PrefixNameSplitSymbol + movement
		muscle := body.Muscle{
			MoveDirection: movement,
			KeyPrefix:     keyPrefix,
		}
		utils.LoadFromMapByKeyPrefix(lb.Organ, keyPrefix, &muscle)
		utils.StoreToMap(lb.Organ, movement, &muscle)
	}
}

func (lb *LeechBody) Action(command string) {

}

func (lb *LeechBody) SenseFromEnv(env interact.Environment) { // get environment info
}

type LeechBrain struct {
	Area *sync.Map
}

func (lb *LeechBrain) InitBrain(wg *sync.WaitGroup, processController *sync.Map) {
	// init sense
	wg.Add(1)
	go func(wg *sync.WaitGroup, processController *sync.Map) {
		defer wg.Done()

		for _, senseNeureType := range config.PrefixSkinAndSenseType {
			for _, senseType := range config.PrefixSenseType {
				for _, position := range config.SkinAndSenseNeurePosition {
					if senseType == "painfulType" && !strings.Contains(senseNeureType, "extremely") {
						// painful sense neure only connect extremely skin neure
						continue
					}
					keyPrefix := "sense" + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + senseType + config.PrefixNameSplitSymbol + position
					sense := brain.Sense{
						SenseNeureType: senseNeureType,
						Position:       position,
						SenseType:      senseType,
						KeyPrefix:      keyPrefix,
					}
					wg.Add(1)
					go sense.InitSense(wg, processController)
				}
			}

		}
	}(wg, processController)
}

func (lb *LeechBrain) LoadBrain(wg *sync.WaitGroup) {
	defer wg.Done()

	for _, senseNeureType := range config.PrefixSkinAndSenseType {
		for _, senseType := range config.PrefixSenseType {
			for _, position := range config.SkinAndSenseNeurePosition {
				if senseType == "painfulType" && !strings.Contains(senseNeureType, "extremely") {
					// painful sense neure only connect extremely skin neure
					continue
				}
				keyPrefix := "sense" + config.PrefixNameSplitSymbol + senseNeureType + config.PrefixNameSplitSymbol + senseType + config.PrefixNameSplitSymbol + position
				sense := brain.Sense{
					SenseNeureType: senseNeureType,
					Position:       position,
					SenseType:      senseType,
					KeyPrefix:      keyPrefix,
				}
				// store four types so that can get one type very fast
				utils.LoadFromMapByKeyPrefix(lb.Area, keyPrefix, &sense)
				utils.StoreToMap(lb.Area, senseNeureType, &sense)
				utils.StoreToMap(lb.Area, position, &sense)
				utils.StoreToMap(lb.Area, senseType, &sense)
			}
		}
	}
}

func (lb *LeechBrain) Sense2Action() (bodyAction string) {
	return
}

type Leech struct {
	Brain *LeechBrain
	Body  *LeechBody
}

func (l *Leech) InitLeech() {
	var wg sync.WaitGroup
	var processController sync.Map

	processController.Store("senseCreateDone", make(chan bool, 1))

	l.Brain = &LeechBrain{}
	l.Body = &LeechBody{}

	// init a leech
	go l.Brain.InitBrain(&wg, &processController)
	go l.Body.InitBody(&wg, &processController)

	wg.Wait()
}

func (l *Leech) LoadLeech() {
	var wg sync.WaitGroup
	wg.Add(2)
	go l.Body.LoadBody(&wg)
	go l.Brain.LoadBrain(&wg)
	wg.Wait()
}

func (l *Leech) Environment2Action(env interact.Environment) string { // get environment param and decide an action
	return ""
}

func (l *Leech) WakeUp() {

}
