package neure

import (
	"encoding/json"
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/utils"
	"time"
)

// Synapse /ˈsɪnæps/ 突触
// Dendrite /ˈdendraɪt/ 樹突
// axon /ˈæksɑːn/ 軸突

type Neure struct {
	Synapses               []*Synapse `json:"sy"` // 軸突連接的突觸，有些神经元有多个突触，一旦激发，所有连接的突触都会尝试激活
	NeureType              string     `json:"nt"`
	NowLinkedDendritesIds  []string   `json:"nld"` // 現在已連接的树突前神经元编号
	ElectricalConductivity int32      `json:"ec"`  // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值，但好像对程序模拟的大脑没什么作用。
	ThisNeureId            string     `json:"tn"`  // the id of database
	NowWeight              float32    `json:"nw"`  // 现在的权重，每刺激一次，增加一点，直到超过weight就被激活，被激活后会reset，超过一段时间无刺激也会reset
	LastTimeActivate       time.Time  `json:"lt"`  // 最后一次激活的时间，精确到纳秒，可以在byte中自由转换
	LastTimeResetNowWeight time.Time  `json:"lw"`  // 最后一次充值now weight的时间
}

func (n *Neure) GetLastTimeActivate() time.Time {
	return n.LastTimeActivate
}

func (n *Neure) TryActivate(weight float32) (activate bool) {
	now := time.Now()
	if now.Sub(n.LastTimeActivate) > config.RefractoryDuration {
		// only activate when neure not in refractory duration
		if now.Sub(n.LastTimeResetNowWeight) > config.RefreshNowWeightDuration {
			// because RefractoryDuration much more small than RefreshNowWeightDuration, so we can put it here
			n.NowWeight = 0
			n.LastTimeResetNowWeight = now
		}
		n.NowWeight += weight
		if n.NowWeight > config.Weight {
			activate = true
			n.NowWeight = 0
			n.LastTimeResetNowWeight = now
			n.LastTimeActivate = now
		}
	}
	return
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

func (n *Neure) GetNeureById(id string) {
	np, ok := NeureMap.Load(id)
	if !ok {
		neureByte := database.GetNeure(id)
		n.Byte2Struct(neureByte)
		NeureMap.Store(id, n)
	} else {
		n = np.(*Neure)
	}
}

func (n *Neure) DeleteNeure() {
	// delete the dendrites of next neures
	for _, synapse := range n.Synapses {
		nextNeure := Neure{}
		nextNeure.GetNeureById(synapse.NextNeureID)
		utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
		nextNeure.UpdateNeure2DB()
	}

	// delete the synapse of pre neures
	for _, dendriteId := range n.NowLinkedDendritesIds {
		preNeure := Neure{}
		preNeure.GetNeureById(dendriteId)
		preNeure.Synapses = RemoveUniqueValueFromSynapse(n.ThisNeureId, preNeure.Synapses)
		preNeure.UpdateNeure2DB()
	}

	_, ok := NeureMap.Load(n.ThisNeureId)
	if ok {
		NeureMap.Delete(n.ThisNeureId)
	}
	// finally, delete this neure
	database.DeleteNeure(n.ThisNeureId)
}

func (n *Neure) DeleteConnection(synapse *Synapse) {
	// remove next neure's dendrites connect
	nextNeure := Neure{}
	nextNeure.GetNeureById(synapse.NextNeureID)
	utils.RemoveUniqueValueFromSlice(n.ThisNeureId, &nextNeure.NowLinkedDendritesIds)
	nextNeure.UpdateNeure2DB()

	// delete from Synapses
	n.Synapses = RemoveUniqueValueFromSynapse(synapse.NextNeureID, n.Synapses)
	n.UpdateNeure2DB()
}

func (n *Neure) ConnectNextNuere(nextNeure *Neure) {
	var synapse Synapse
	synapse.NextNeureID = nextNeure.ThisNeureId
	n.Synapses = append(n.Synapses, &synapse)
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
