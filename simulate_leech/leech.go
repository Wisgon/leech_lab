package leech

import (
	"graph_robot/interact"
	"graph_robot/neure"
)

type LeechBody struct {
	Skin     []*neure.Neure // sensory organ of leech
	Nose     []*neure.Neure
	Movement []*neure.Neure
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
	go l.brain.InitBrain()
	go l.body.InitBody()
}

func (l *Leech) Environment2Action(env interact.Environment) string { // get environment param and decide an action
	return ""
}
