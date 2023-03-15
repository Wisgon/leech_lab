package neure

import (
	"encoding/json"
	"graph_robot/database"
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
	NextNeureID int64 `json:"n1"` // 突觸後神經元，是這個軸突所連接的神經元
	Weight      int32 `json:"iw"` // 與nextNeure的連接權重
}

type Neure struct {
	AxonSynapse Synapse `json:"sa"` // 軸突連接的突觸，有些神经元有多个突触，但是现在还未明白多个或单个突触有什么影响
	// dendrites number should be infinite, so next line is commented
	// DendritesLinkNum       int32   `json:"ld"`  // 樹突的數量
	NowLinkedDendritesNum  int32 `json:"ndn"` // 現在已連接的樹突的數量
	NeureType              bool  `json:"tn"`  // true為激發神經元，false為抑制神經元
	ElectricalConductivity int32 `json:"ce"`  // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值
	ThisNeureId            int64 `json:"did"` // the id of database
}

// func (n *Neure) IncreaseDendritesNum() {
// 	// 神經元的樹突與其他神經元的軸突連接時要加1，返回連接成功或失敗的結果
// 	n.NowLinkedDendritesNum += 1
// }

func (n *Neure) CreateNeureInDB() {
	databaseModel := database.NeureData{Neure: []byte{}}
	id := databaseModel.Create()
	n.ThisNeureId = id
	// to update neure because the neure byte of databaseModel which created in first step is empty, n.ThisNeureId is 0
	n.SaveNeure2DB()
}

func (n *Neure) SaveNeure2DB() {
	databaseModel := database.NeureData{
		ID:     n.ThisNeureId,
		Neure:  n.Struct2Byte(),
		Linked: n.AxonSynapse.NextNeureID != 0,
	}
	databaseModel.Save()
}

func (n *Neure) GetNeureFromDatabaseById(id int64) {
	databaseModel := database.NeureData{}
	databaseModel.GetNeureDataById(id)
	n.Byte2Struct(databaseModel.Neure)
}

func (n *Neure) ConnectNextNuere(nextId int64) {
	n.AxonSynapse.NextNeureID = nextId
	n.SaveNeure2DB()
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
