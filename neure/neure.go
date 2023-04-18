package neure

import (
	"encoding/json"
	"fmt"
	"graph_robot/config"
	"graph_robot/creature"
	"graph_robot/database"
	"graph_robot/utils"
	"sync"
)

// Synapse /ˈsɪnæps/ 突触
// Dendrite /ˈdendraɪt/ 樹突
// axon /ˈæksɑːn/ 軸突

type Synapse struct {
	// 突觸，連接兩個Neure
	NextNeureID string `json:"n1"` // 突觸後神經元，是這個軸突所連接的神經元
	SynapseNum  int32  `json:"uy"` // 连接到next neure的突触数量，跟长时记忆有关，长时记忆的连接突触数量会变多
	Status      string `json:"us"` // 该突触所处的状态，如强化态，弱化态，生长态等，生长态会持续一段较长的时间
}

func (s *Synapse) GetNextId() string {
	return s.NextNeureID
} // use to fit an interface

func (s *Synapse) SetNextId(nextNeureId string) {
	s.NextNeureID = nextNeureId
}

func (s *Synapse) CheckLinkStrength() {
	// 设计一个函数，连接强度越小越容易被清除，强度越大越难清除
}

type NormalSynapse struct {
	Synapse
}

func (s *NormalSynapse) ActivateNextNeure() {
	// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
}

type RegulateSynapse struct {
	// 这是调节神经元的突触，不同类型的突触有不同的ActivateNextNeure方法
	Synapse
}

func (s *RegulateSynapse) ActivateNextNeure() {
	// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
}

type Neure struct {
	mu                     sync.Mutex
	Synapses               []creature.Synapse `json:"sa"`  // 軸突連接的突觸，有些神经元有多个突触，但是现在还未明白多个或单个突触有什么影响，一旦激发，所有连接的突触都会激活
	NowLinkedDendritesIds  []string           `json:"ndn"` // 現在已連接的樹突
	ElectricalConductivity int32              `json:"ce"`  // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值，但好像对程序模拟的大脑没什么作用。
	ThisNeureId            string             `json:"did"` // the id of database
	Weight                 float32            `json:"iw"`  //触发这个神经元产生动作电位的weight
	NowWeight              float32            `json:"tw"`  // 现在的权重，每刺激一次，增加一点，直到超过weight就被激活
	NeureType              string             `json:"tn"`
	// RefractoryPeriod float64 `json:"ir"`  //不应期，不过是否有必要还有待商榷
}

func (n *Neure) AddNowWeight(weight float32) (activate bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.NowWeight += weight
	if n.NowWeight > n.Weight {
		activate = true
		n.ResetNowWeight()
	}
	return
}

func (n *Neure) ResetNowWeight() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.NowWeight = 0
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
	for _, synapse := range n.Synapses {
		nextNeure := Neure{}
		nextNeure.GetNeureFromDbById(synapse.GetNextId())
		utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
		nextNeure.UpdateNeure2DB()
	}

	// delete the synapse of pre neures
	for _, dendriteId := range n.NowLinkedDendritesIds {
		preNeure := Neure{}
		preNeure.GetNeureFromDbById(dendriteId)
		preNeure.Synapses = utils.RemoveUniqueValueFromSynapse(n.ThisNeureId, preNeure.Synapses)
		preNeure.UpdateNeure2DB()
	}

	// finally, delete this neure
	database.DeleteNeure(n.ThisNeureId)
}

func (n *Neure) DeleteConnection(synapse *Synapse) {
	// remove next neure's dendrites connect
	nextNeure := Neure{}
	nextNeure.GetNeureFromDbById(synapse.NextNeureID)
	utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
	nextNeure.UpdateNeure2DB()

	// delete from Synapses
	n.Synapses = utils.RemoveUniqueValueFromSynapse(synapse.NextNeureID, n.Synapses)
	n.UpdateNeure2DB()
}

func (n *Neure) ConnectNextNuere(nextNeure *Neure) {
	var synapse creature.Synapse
	if n.NeureType == "regulate" {
		synapse = &RegulateSynapse{}
	} else if n.NeureType == "normal" {
		synapse = &NormalSynapse{}
	} else {
		panic("neure type error!")
	}
	synapse.SetNextId(nextNeure.ThisNeureId)
	n.Synapses = append(n.Synapses, synapse)
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
