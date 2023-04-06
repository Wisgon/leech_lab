package creature

type Brain interface {
	InitBrain()                        // 建造初始神经元和初始entrance
	Sense2Action() (bodyAction string) // get sense from sensory organ and output an action
}
