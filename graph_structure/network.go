package graph_structure

import (
	"graph_robot/database"
)

type NetWork struct {
	Neures map[int64]*Neure
}

func (nw *NetWork) LoadNetwork(neureId int64) {
	neure := database.GetNeureById(neureId)
	nw.Neures[neureId] = Byte2Struct(neure.Neure)
	// todo: 這裡肯定不是簡單的為0就加載完了，肯定有其他條件，以後想到再加
	if nw.Neures[neureId].AxonSynapse.NextNeureID != 0 {
		nw.LoadNetwork(nw.Neures[neureId].AxonSynapse.NextNeureID)
	}
}

func (nw *NetWork) SaveNetwork() {
	for id, value := range nw.Neures {
		var neure = database.Neures{
			ID:    id,
			Neure: Struct2Byte(value),
		}
		database.UpdateNeures([]database.Neures{neure})
	}
}
