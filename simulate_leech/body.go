package leech

import (
	"graph_robot/interact"
	"graph_robot/neure"
)

type LeechBody struct {
	Skin []neure.Neure // sensory organ of leech
}

func (b *LeechBody) Action(command string) {

}

func (b *LeechBody) Sense(env interact.Environment) { // get environment info

}
