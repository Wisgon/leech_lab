package brain

import "graph_robot/interact"

type Brain interface {
	InitBrain()                                  // 建造初始神经元和初始entrance
	Environment2Action(env interact.Environment) // get environment param and decide an action
}
