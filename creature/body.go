package creature

import (
	"graph_robot/interact"
)

type Body interface {
	Sense(env interact.Environment) // get environment info
	Action(command string)          // act an action to world and get response
}
