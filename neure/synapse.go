package neure

import (
	"graph_robot/config"
	"math/rand"
	"sync"
)

type Synapse struct {
	// 突觸，連接兩個Neure
	mu           sync.Mutex
	NextNeureID  string  `json:"i"` // 突觸後神經元，是這個軸突所連接的神經元
	SynapseNum   int32   `json:"j"` // 连接到next neure的突触数量，跟长时记忆有关，长时记忆的连接突触数量会变多，初始时必须最少为1
	LinkStrength float32 `json:"k"` // 连接强度，在长时程增强的时候增强，过后减弱
}

func (s *Synapse) CheckLinkStrength() {
	// 设计一个函数，连接强度越小越容易被清除，强度越大越难清除
}

func (s *Synapse) ActivateNextNeure(neureType string) (ok bool, nextNeure *Neure) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch neureType {
	case "normal":
		// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
		nextNeure = GetNeureById(s.NextNeureID)
		ok = nextNeure.TryActivate(s.LinkStrength * float32(s.SynapseNum))
		if !ok {
			if (1-nextNeure.NowWeight/config.Weight)*config.BreakThroughCoefficient > rand.Float32() {
				// 再增加一次now weight的概率越大，但最大不会超过设置值（如0.3）
				ok = nextNeure.TryActivate(s.LinkStrength * float32(s.SynapseNum))
			}
		}
	case "regulate":
		// 这是调节神经元的突触，不同类型的突触有不同的ActivateNextNeure方法
	case "inhibitory":
		// 抑制型神经元
	default:
		panic("neure type worng!!!")
	}
	return
}
