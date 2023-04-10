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
	lb.initSkin()
	lb.initNose()
	lb.initMovement()
}

func (lb *LeechBody) initSkin() {

}

func (lb *LeechBody) initNose() {

}

func (lb *LeechBody) initMovement() {

}

func (lb *LeechBody) Action(command string) {

}

func (lb *LeechBody) Sense(env interact.Environment) { // get environment info
}
