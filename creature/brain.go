package creature

import "graph_robot/neure"

type Brain interface {
	InitBrain()                                                 // 建造初始神经元和初始entrance
	Sense2Action(neureResult []neure.Neure) (bodyAction string) // get sense from sensory organ and output an action
}
