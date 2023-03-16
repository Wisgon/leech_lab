package leech

import (
	"graph_robot/interact"
	"graph_robot/neure"
)

type LeechBrain struct {
	SensoryOrgan neure.NeureEntrance // sensory organ is use for input the env signal
}

func (lb *LeechBrain) InitBrain() {
}

func (lb *LeechBrain) Environment2Action(env interact.Environment) string {
	return ""
}
