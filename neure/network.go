package neure

// notice: NetWork was abandoned

// import (
// 	"graph_robot/database"
// )

//小知识：
// 如果神经冲动的传导是从树突传到轴突的话，那么一个network一定是一颗倒过来的数，神经冲动的入口是树枝，最后传导到顶端某些终点作为输出

// type NetWork struct {
// 	Neures           map[int64]*Neure
// 	NeureOrder       []int64
// 	NeedUpdateNeures []int64
// }

// func (nw *NetWork) LoadNetwork(neureId int64) {
// 	neure := database.GetNeureById(neureId)
// 	nw.Neures[neureId] = Byte2Struct(neure.Neure)
// 	nw.NeureOrder = append(nw.NeureOrder, neureId)
// 	if nw.Neures[neureId].AxonSynapse.NextNeureID != 0 {
// 		nw.LoadNetwork(nw.Neures[neureId].AxonSynapse.NextNeureID)
// 	}
// }

// func (nw *NetWork) SaveNetwork() {
// 	var needUpdateNeures = make(map[int64]*Neure)
// 	for _, index := range nw.NeedUpdateNeures {
// 		needUpdateNeures[index] = nw.Neures[index]
// 	}
// 	for id, neure := range needUpdateNeures {
// 		var neureDb = database.NeureDb{
// 			ID:    id,
// 			Neure: Struct2Byte(neure),
// 		}
// 		database.UpdateNeures([]database.NeureDb{neureDb})
// 	}
// }
