package leech

import (
	"graph_robot/neure"
)

type LeechBrain struct {
	SensoryEntrance neure.NeureEntrance // sensory entrance is use for get the signal from sensory organ
}

func (lb *LeechBrain) InitBrain() {
}

func (lb *LeechBrain) Sense2Action(neureResult []neure.Neure) (bodyAction string) {
	return ""
}
