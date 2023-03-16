package leech

import "graph_robot/interact"

type Leech struct {
	Brain LeechBrain
	Body  LeechBody
}

func (l *Leech) Environment2Action(env interact.Environment) string { // get environment param and decide an action
	return ""
}
