package neure

import (
	"graph_robot/config"
	"math/rand"
)

type Synapse struct {
	// 突觸，連接兩個Neure
	NextNeureID  string  `json:"n1"` // 突觸後神經元，是這個軸突所連接的神經元
	SynapseNum   int32   `json:"uy"` // 连接到next neure的突触数量，跟长时记忆有关，长时记忆的连接突触数量会变多，初始时必须最少为1
	LinkStrength float32 `json:"gk"` // 连接强度，在长时程增强的时候增强，过后减弱
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

func (s *NormalSynapse) ActivateNextNeure() (ok bool, nextNeure *Neure[*NormalSynapse]) {
	// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
	nextNeure = &Neure[*NormalSynapse]{}
	nextNeure.GetNeureFromDbById(s.NextNeureID)
	ok = nextNeure.AddNowWeight(s.LinkStrength * float32(s.SynapseNum))
	if !ok {
		if (1-nextNeure.NowWeight/config.Weight)*config.BreakThroughCoefficient > rand.Float32() {
			// 再增加一次now weight的概率越大，但最大不会超过设置值（如0.3）
			ok = nextNeure.AddNowWeight(s.LinkStrength * float32(s.SynapseNum))
		}
	}
	return
}

type RegulateSynapse struct {
	// 这是调节神经元的突触，不同类型的突触有不同的ActivateNextNeure方法
	Synapse
}

func (s *RegulateSynapse) ActivateNextNeure() (ok bool, nextNeure *Neure[*RegulateSynapse]) {
	// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
	return
}

type InhibitorySynapse struct {
	//抑制神经元
	Synapse
}

func (s *InhibitorySynapse) ActivateNextNeure() (ok bool, nextNeure *Neure[*InhibitorySynapse]) {
	// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为，抑制神经元让下一个now weight变负数
	return
}
