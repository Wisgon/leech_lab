package leech

import (
	"graph_robot/interact"
	"graph_robot/simulate_leech/body"
)

type LeechBody struct {
	Skin     body.Skin // sensory organ of leech
	Movement body.Movement
}

func (lb *LeechBody) InitBody() {
}

func (lb *LeechBody) Action(command string) {

}

func (lb *LeechBody) Sense(env interact.Environment) { // get environment info
}

type LeechBrain struct {
}

func (lb *LeechBrain) InitBrain() {

}

func (lb *LeechBrain) Sense2Action() (bodyAction string) {
	return
}

type Leech struct {
	brain LeechBrain
	body  LeechBody
}

func (l *Leech) InitLeech() {
	// init a leech
	l.brain.InitBrain()
	l.body.InitBody()
}

func (l *Leech) Environment2Action(env interact.Environment) string { // get environment param and decide an action
	return ""
}
