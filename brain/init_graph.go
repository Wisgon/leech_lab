package brain

import (
	gs "graph_robot/graph_structure"
)

func InitBrainGraph() {
	//todo:初始化网络，这个网络要有很多个分区，每个分区都有很多network，每个分区神经元network可无限增长。
	var eyes = gs.NeureEntrance{}
	eyes.GetNeureEntranceFromDbByTypeName("eyes")
	var brain = Brain{
		Eyes: eyes,
	}

	go brain.Think()
}
