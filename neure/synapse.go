package neure

import (
	"graph_robot/config"
	"log"
	"math"
	"sync"
	"time"
)

type Synapse struct {
	// 突觸，連接兩個Neure
	mu                      sync.Mutex
	NextNeureID             string  `json:"i"` // 突觸後神經元，是這個軸突所連接的神經元
	SynapseNum              int32   `json:"j"` // 连接到next neure的突触数量，跟长时记忆有关，长时记忆的连接突触数量会变多，初始时必须最少为1
	LinkStrength            float32 `json:"k"` // 连接强度，在长时程增强的时候增强，过后减弱
	NextNeureSynapseId      string  `json:"l"` // 这个只有regulate和inhibitory神经元才有的，方便找到下一个调节的synapse
	Hibituationbility       bool    `json:"m"` // not all synapse need hibituation
	AttenuationAccumulative float64 `json:"n"` // 衰减量总计,如果超过1，则向下取整化为整数然后从突出数量中扣掉
}

func (s *Synapse) ActivateNextNeure(neureType string) (nextNeure *Neure) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch neureType {
	case config.PrefixNeureType["common"]:
		// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
		nextNeure = GetNeureById(s.NextNeureID)
		// send signal, try to activate next neure, but we can't know whether it is activate, only next neure know.
		nextNeure.SignalChannel <- s.LinkStrength * float32(s.SynapseNum)
		if s.Hibituationbility && s.SynapseNum > 1 {
			// each time signal comes, and then there is no synapse enhance, SynapseNum will reduce by a Attenuation Function because of the hibituation
			s.AttenuationAccumulative += AttenuationFunction(s.SynapseNum)
			// when s.SynapseNum is very big, AttenuationFunction(s.SynapseNum) is very small. so the more synapseNum is, the stable the link is.
			if s.AttenuationAccumulative > 1 {
				// when AttenuationAccumulative is bigger than 1, it can be turn to int so that SynapseNum can subtract it.
				s.SynapseNum -= int32(s.AttenuationAccumulative)
				if s.SynapseNum < 1 {
					// at least there is 1 Synapse, when the SynapseNum is 1, it is almost unactivatable
					s.SynapseNum = 1
				}
				s.AttenuationAccumulative = 0
			}
		}
	case config.PrefixNeureType["regulate"]:
		// 这是调节神经元的突触，不同类型的突触有不同的ActivateNextNeure方法
		if nextNeure.NeureType == config.PrefixNeureType["common"] {
			// regulate won't activate next neure if next neure is common neure, it will regulate the linkstrength of next neure
		} else {
			nextNeure.SignalChannel <- config.Weight + 1 // directly activate
		}
	case config.PrefixNeureType["inhibitory"]:
		// 抑制型神经元
	default:
		log.Panic("neure type worng!!!")
	}
	return
}

func (s *Synapse) ActivateAtOneFrequency(neureType string, activateDuration int, activateTimes int) {
	// activateDuration是隔多少秒刺激一次，activateTimes是总共要激活多少次
	for i := 0; i < activateTimes; i++ {
		s.ActivateNextNeure(neureType)
		time.Sleep(time.Duration(activateDuration) * time.Second)
	}
}

func AttenuationFunction(synapseNum int32) float64 {
	// when synapseNum is very big, result is very small.
	return math.Pow(0.5, 0.05*float64((int(synapseNum)-config.AttenuationFunctionFactor)))
}
