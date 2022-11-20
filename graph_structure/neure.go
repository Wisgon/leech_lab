package graph_structure

import (
	"graph_robot/utils"
	"sync"
	"time"
)

// Synapse /ˈsɪnæps/ 突触
// Dendrite /ˈdendraɪt/ 樹突
// axon /ˈæksɑːn/ 軸突

// 小知識：
// 1. 神經元神經衝動傳導方向是樹突傳向軸突然後傳到下一個神經元的樹突
// 2. 一個神經元的軸突只能連接一個神經元的樹突或細胞體，但是樹突可以連接多個神經元，將這些神經元的信號通過軸突傳給下一個神經元
// 3. 神經元現在海馬體形成短時記憶，然後再在皮質層形成長期記憶

var Synapses sync.Map

type Synapse struct {
	// 突觸，連接兩個Neure
	NextNeureID string // 突觸後神經元，是提供樹突的神經元
	PreNeureID  string // 突出前神經元，是提供軸突的神經元
}

type Neure struct {
	AxonSynapseID          string // 軸突連接的突觸
	Weight                 int    // 與nextNeure的連接權重
	DendritesLinkNum       int    // 樹突的數量
	NowLinkedDendritesNum  int    // 現在已連接的樹突的數量
	NeureType              bool   // true為激發神經元，false為抑制神經元
	ElectricalConductivity int    // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值
	HashId                 string // 在整個graph的map中的編號
}

func (n *Neure) IncreaseDendritesNum() {
	// 神經元的樹突與其他神經元的軸突連接時要加1，返回連接成功或失敗的結果
	n.NowLinkedDendritesNum += 1
}

func (n *Neure) DeleteLink() {
	Synapses.Delete(n.AxonSynapseID)
	n.AxonSynapseID = ""
}

func (n *Neure) LinkNextNeure(nextNeure *Neure, weight int) (linkSuccessed bool) {
	if nextNeure.NowLinkedDendritesNum < nextNeure.DendritesLinkNum {
		var snp Synapse
		synapseID := utils.GetUniqueId(time.Now().UnixNano())
		snp.PreNeureID = n.HashId
		snp.NextNeureID = nextNeure.HashId
		n.AxonSynapseID = synapseID
		n.Weight = weight
		nextNeure.IncreaseDendritesNum()
		Synapses.Store(synapseID, snp)
		return true
	}
	return false
}
