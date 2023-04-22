package neure

import (
	"encoding/json"
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/utils"
	"sync"
)

// Synapse /ˈsɪnæps/ 突触
// Dendrite /ˈdendraɪt/ 樹突
// axon /ˈæksɑːn/ 軸突
type ST interface {
	*NormalSynapse | *InhibitorySynapse | *RegulateSynapse
	GetNextId() string
	SetNextId(string)
}

type Neure[T ST] struct {
	mu                     sync.Mutex
	Synapses               []T      `json:"sa"`  // 軸突連接的突觸，有些神经元有多个突触，但是现在还未明白多个或单个突触有什么影响，一旦激发，所有连接的突触都会激活
	NowLinkedDendritesIds  []string `json:"ndn"` // 現在已連接的樹突
	ElectricalConductivity int32    `json:"ce"`  // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值，但好像对程序模拟的大脑没什么作用。
	ThisNeureId            string   `json:"did"` // the id of database
	NowWeight              float32  `json:"tw"`  // 现在的权重，每刺激一次，增加一点，直到超过weight就被激活，被激活后会reset，超过一段时间无刺激也会reset
	LastTimeActivate       float32  `json:"ii"`  // 最后一次激活的时间，单位为纳秒
}

func (n *Neure[T]) GetSynapses() []T {
	return n.Synapses
}

func (n *Neure[T]) AddNowWeight(weight float32) (activate bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.NowWeight += weight
	if n.NowWeight > config.Weight {
		activate = true
		n.ResetNowWeight()
	}
	return
}

func (n *Neure[T]) ResetNowWeight() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.NowWeight = 0
}

func (n *Neure[T]) CreateNeureInDB(keyPrefix string) {
	uniqueNum := database.GetSeqNum(keyPrefix)
	key := keyPrefix + config.PrefixNumSplitSymbol + fmt.Sprint(uniqueNum)
	n.ThisNeureId = key
	database.CreateNeure(n.Struct2Byte(), key)
}

func (n *Neure[T]) UpdateNeure2DB() {
	database.UpdateNeure(n.Struct2Byte(), n.ThisNeureId)
}

func (n *Neure[T]) GetNeureFromDbById(id string) {
	neureByte := database.GetNeure(id)
	n.Byte2Struct(neureByte)
}

func (n *Neure[T]) DeleteNeure() {
	// delete the dendrites of next neures
	for _, synapse := range n.Synapses {
		nextNeure := Neure[T]{}
		nextNeure.GetNeureFromDbById(synapse.GetNextId())
		utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
		nextNeure.UpdateNeure2DB()
	}

	// delete the synapse of pre neures
	for _, dendriteId := range n.NowLinkedDendritesIds {
		preNeure := Neure[T]{}
		preNeure.GetNeureFromDbById(dendriteId)
		preNeure.Synapses = RemoveUniqueValueFromSynapse(n.ThisNeureId, preNeure.Synapses)
		preNeure.UpdateNeure2DB()
	}

	// finally, delete this neure
	database.DeleteNeure(n.ThisNeureId)
}

func (n *Neure[T]) DeleteConnection(synapse *Synapse) {
	// remove next neure's dendrites connect
	nextNeure := Neure[T]{}
	nextNeure.GetNeureFromDbById(synapse.NextNeureID)
	utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
	nextNeure.UpdateNeure2DB()

	// delete from Synapses
	n.Synapses = RemoveUniqueValueFromSynapse(synapse.NextNeureID, n.Synapses)
	n.UpdateNeure2DB()
}

func (n *Neure[T]) ConnectNextNuere(nextNeure *Neure[T]) {
	var synapse T
	synapse.SetNextId(nextNeure.ThisNeureId)
	n.Synapses = append(n.Synapses, synapse)
	n.UpdateNeure2DB()
	nextNeure.NowLinkedDendritesIds = append(nextNeure.NowLinkedDendritesIds, n.ThisNeureId) // next neure dendrites append
	nextNeure.UpdateNeure2DB()
}

func (n *Neure[T]) Struct2Byte() []byte {
	nb, err := json.Marshal(n)
	if err != nil {
		panic("json marshal error: " + err.Error())
	}
	return nb
}

func (n *Neure[T]) Byte2Struct(neureByte []byte) {
	err := json.Unmarshal(neureByte, n)
	if err != nil {
		panic("json unmarshal error: " + err.Error())
	}
}
