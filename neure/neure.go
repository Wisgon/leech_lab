package neure

import (
	"encoding/json"
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/utils"
)

// Synapse /ˈsɪnæps/ 突触
// Dendrite /ˈdendraɪt/ 樹突
// axon /ˈæksɑːn/ 軸突

// 小知識：
// 1. 神經元神經衝動傳導方向是樹突傳向軸突然後傳到下一個神經元的樹突
// 2. 一個神經元的軸突只能連接一個神經元的樹突或細胞體，但是樹突可以連接多個神經元，將這些神經元的信號通過軸突傳給下一個神經元
// 3. 神經元現在海馬體形成短時記憶，然後再在皮質層形成長期記憶

type Synapse struct {
	// 突觸，連接兩個Neure
	NextNeureID string `json:"n1"` // 突觸後神經元，是這個軸突所連接的神經元
	Weight      int32  `json:"iw"` // 與nextNeure的連接權重
	NowWeight   int32  `json:"tw"` // 现在的权重，每刺激一次，增加一点，直到超过weight就被激活
}

func (s Synapse) GetNextId() string {
	return s.NextNeureID
} // use to fit an interface

type Neure struct {
	AxonSynapse           []Synapse `json:"sa"`  // 軸突連接的突觸，有些神经元有多个突触，但是现在还未明白多个或单个突触有什么影响
	NowLinkedDendritesIds []string  `json:"ndn"` // 現在已連接的樹突
	// NeureType              bool     `json:"tn"`  // true為激發神經元，false為抑制神經元
	ElectricalConductivity int32   `json:"ce"`  // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值
	ThisNeureId            string  `json:"did"` // the id of database
	RefractoryPeriod       float64 `json:"ir"`  //不应期，不过是否有必要还有待商榷
}

func (n *Neure) CreateNeureInDB(keyPrefix string) {
	uniqueNum := database.GetSeqNum(keyPrefix)
	key := keyPrefix + config.PrefixNumSplitSymbol + fmt.Sprint(uniqueNum)
	n.ThisNeureId = key
	database.CreateNeure(n.Struct2Byte(), key)
}

func (n *Neure) UpdateNeure2DB() {
	database.UpdateNeure(n.Struct2Byte(), n.ThisNeureId)
}

func (n *Neure) GetNeureFromDbById(id string) {
	neureByte := database.GetNeure(id)
	n.Byte2Struct(neureByte)
}

func (n *Neure) DeleteNeure() {
	// delete the dendrites of next neures
	for _, synapse := range n.AxonSynapse {
		nextNeure := Neure{}
		nextNeure.GetNeureFromDbById(synapse.NextNeureID)
		utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
		nextNeure.UpdateNeure2DB()
	}

	// delete the synapse of pre neures
	for _, dendriteId := range n.NowLinkedDendritesIds {
		preNeure := Neure{}
		preNeure.GetNeureFromDbById(dendriteId)
		utils.RemoveUniqueValueFromSynapse(n.ThisNeureId, &preNeure.AxonSynapse)
		preNeure.UpdateNeure2DB()
	}

	// finally, delete this neure
	database.DeleteNeure(n.ThisNeureId)
}

func (n *Neure) ConnectNextNuere(nextNeure *Neure) {
	synapse := Synapse{}
	synapse.NextNeureID = nextNeure.ThisNeureId
	n.AxonSynapse = append(n.AxonSynapse, synapse)
	n.UpdateNeure2DB()
	nextNeure.NowLinkedDendritesIds = append(nextNeure.NowLinkedDendritesIds, n.ThisNeureId) // next neure dendrites append
	nextNeure.UpdateNeure2DB()
}

func (n *Neure) Struct2Byte() []byte {
	nb, err := json.Marshal(n)
	if err != nil {
		panic("json marshal error: " + err.Error())
	}
	return nb
}

func (n *Neure) Byte2Struct(neureByte []byte) {
	err := json.Unmarshal(neureByte, n)
	if err != nil {
		panic("json unmarshal error: " + err.Error())
	}
}
