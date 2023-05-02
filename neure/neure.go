package neure

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/database"
	"log"
	"sync"
	"time"
)

// Synapse /ˈsɪnæps/ 突触
// Dendrite /ˈdendraɪt/ 樹突
// axon /ˈæksɑːn/ 軸突

type Neure struct {
	mu                     sync.Mutex
	Synapses               map[string]*Synapse `json:"a"` // 軸突連接的突觸，有些神经元有多个突触，一旦激发，所有连接的突触都会尝试激活
	NeureType              string              `json:"b"` //神经元的类型，有普通神经元，调节神经元和抑制神经元
	NowLinkedDendritesIds  map[string]struct{} `json:"c"` // 現在已連接的树突前神经元编号
	ElectricalConductivity int32               `json:"d"` // 導電性，越大這個軸突導電性越弱，因為每次經過這個軸突，電流強度都要減去這個值，但好像对程序模拟的大脑没什么作用。
	ThisNeureId            string              `json:"e"` // the id of database
	NowWeight              float32             `json:"f"` // 现在的权重，每刺激一次，增加一点，直到超过weight就被激活，被激活后会reset，超过一段时间无刺激也会reset
	LastTimeActivate       time.Time           `json:"g"` // 最后一次激活的时间，精确到纳秒，可以在byte中自由转换
	LastTimeResetNowWeight time.Time           `json:"h"` // 最后一次重置now weight的时间
}

func (n *Neure) SaveNeure2Db() {
	database.CreateData(n.Struct2Byte(), n.ThisNeureId)
}

func (n *Neure) UpdateNeure2DB() {
	// 这里不能用mutex锁， 不然在neure map的range中执行update会报错
	database.UpdateNeure(n.Struct2Byte(), n.ThisNeureId)
}

func (n *Neure) AddNowDendrites(preNeureId string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.NowLinkedDendritesIds[preNeureId] = struct{}{}
}

func (n *Neure) RemoveDendrites(preNeureId string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.NowLinkedDendritesIds, preNeureId)
}

func (n *Neure) ChangeElectricalConductivity(value int, op string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	switch op {
	case "add":
		n.ElectricalConductivity += int32(value)
	case "sub":
		n.ElectricalConductivity -= int32(value)
	default:
		log.Panic("invalid op with ElectricalConductivity")
	}
}

func (n *Neure) TryActivate(weight float32) (activate bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
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

func (n *Neure) DeleteConnection(synapse *Synapse) {
	n.mu.Lock()
	defer n.mu.Unlock()

	// remove next neure's dendrites connect
	nextNeure := GetNeureById(synapse.NextNeureID)
	nextNeure.RemoveDendrites(n.ThisNeureId)
	// delete from Synapses
	delete(n.Synapses, synapse.NextNeureID)
}

func (n *Neure) ConnectNextNuere(synapse *Synapse) {
	n.mu.Lock()
	defer n.mu.Unlock()

	nextNeure := GetNeureById(synapse.NextNeureID)
	n.Synapses[synapse.NextNeureID] = synapse
	nextNeure.AddNowDendrites(n.ThisNeureId) // next neure dendrites append
}

func (n *Neure) Struct2Byte() []byte {
	n.mu.Lock()
	defer n.mu.Unlock()

	nb, err := json.Marshal(n)
	if err != nil {
		log.Panic("json marshal error: " + err.Error())
	}
	return nb
}

func (n *Neure) Byte2Struct(neureByte []byte) {
	err := json.Unmarshal(neureByte, n)
	if err != nil {
		log.Panic("json unmarshal error: " + err.Error())
	}
}
