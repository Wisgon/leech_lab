package graph_structure

import (
	"graph_robot/database"
)

type NetWork struct {
	Neures           map[int64]*Neure
	NeureOrder       []int64
	NeedUpdateNeures []int64
}

func (nw *NetWork) LoadNetwork(neureId int64) {
	neure := database.GetNeureById(neureId)
	nw.Neures[neureId] = Byte2Struct(neure.Neure)
	nw.NeureOrder = append(nw.NeureOrder, neureId)
	// todo: 這裡肯定不是簡單的為0就加載完了，肯定有其他條件，以後想到再加
	if nw.Neures[neureId].AxonSynapse.NextNeureID != 0 {
		nw.LoadNetwork(nw.Neures[neureId].AxonSynapse.NextNeureID)
	}
}

func (nw *NetWork) SaveNetwork() {
	var needUpdateNeures = make(map[int64]*Neure)
	for _, index := range nw.NeedUpdateNeures {
		needUpdateNeures[index] = nw.Neures[index]
	}
	for id, neure := range needUpdateNeures {
		var neure = database.Neures{
			ID:    id,
			Neure: Struct2Byte(neure),
		}
		database.UpdateNeures([]database.Neures{neure})
	}
}
