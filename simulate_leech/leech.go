package leech

import "graph_robot/interact"

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
