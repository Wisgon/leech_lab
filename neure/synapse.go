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
	NextNeureID             string    `json:"i"` // 突觸後神經元，是這個軸突所連接的神經元
	SynapseNum              int32     `json:"j"` // 连接到next neure的突触数量，跟长时记忆有关，长时记忆的连接突触数量会变多，初始时必须最少为1
	LinkStrength            float64   `json:"k"` // 连接强度，在长时程增强的时候增强，过后减弱
	NextNeureSynapseId      string    `json:"l"` // 这个只有regulate和inhibitory神经元才有的，方便找到下一个调节的synapse
	Hibituationbility       bool      `json:"m"` // not all synapse need hibituation
	AttenuationAccumulation float64   `json:"n"` // 衰减量总计,如果超过1，则向下取整化为整数然后从突触数量中扣掉
	ThisNeureId             string    `json:"o"` // 该突触所在神经元
	LTPAccumulation         float64   `json:"p"` // 长时程增强积累值
	LTPStartTime            time.Time `json:"q"` // 长时程增强的开始时间，这个和持续时间有关，如果超过持续时间，长时程增强结束
	LTPStrength             float64   `json:"u"` // 长时程增强信号的大小，这个在regulate和inhibitory神经元是用来发送的，而在common神经元中，是用来加到LinkStrength上的
}

func (s *Synapse) HandleRegulationStimulate(LTPStrength float64, thisNeure *Neure) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	if now.Sub(s.LTPStartTime) > config.LinkStrengthIncDuration {
		// means that this is a new LTP, in one LinkStrengthIncDuration only strengthen once
		s.LTPStartTime = now
		s.LTPStrength = LTPStrength

		// 设计一个函数，当长时程增强的刺激来到时，可以把最近的activate的神经元事件关联起来，时间越近，那么积累的增强系数就越大，在反复刺激时增强的synapse num就越多
		// 长时程增强的积累值的计算是，新积累值等于新来的LTPStrength加上旧积累值乘于一个时间衰减函数，这个时间是最后刺激的时间，时间越长，衰减值越多，时间越短，越无限接近于1
		durationTime := now.Sub(thisNeure.LastTimeActivate)
		s.LTPAccumulation = LTPStrength + s.LTPAccumulation*LTPAccumulationAttenuationFunction(durationTime.Seconds())

		if s.LTPAccumulation > float64(config.LongTermMemoryLTPAThreshold) {
			s.SynapseNum += int32(config.SurpassThresholdSynapseNumAdd)
			log.Println("debug: neure:", thisNeure.ThisNeureId, " next:", s.NextNeureID, " adding syncNum:", s.SynapseNum)
			s.LTPAccumulation = 0 // reset
		}
	}
}

func (s *Synapse) ActivateNextNeure(thisNeure *Neure) (nextNeure *Neure) {
	s.mu.Lock()
	defer s.mu.Unlock()

	nextNeure = GetNeureById(s.NextNeureID)

	switch thisNeure.NeureType {
	case config.PrefixNeureType["common"]:
		// 激活下一个神经元，根据不同的连接强度和下一个神经元的weight做出不同的行为
		addWeight := s.LinkStrength * float64(s.SynapseNum)
		if time.Since(s.LTPStartTime) < config.LinkStrengthIncDuration {
			// the synapse is in LTP
			addWeight += float64(s.LTPStrength)
		}
		signalInfo := make(map[string]interface{})
		signalInfo["weight"] = addWeight
		signalInfo["source_neure_id"] = thisNeure.ThisNeureId
		// send signal, try to activate next neure, but we can't know whether it is activate, only next neure know.
		go func() {
			nextNeure.SignalChannel <- signalInfo
		}()
		if s.Hibituationbility && s.SynapseNum > 1 {
			// each time signal comes, and then there is no synapse enhance, SynapseNum will reduce by a Attenuation Function because of the hibituation
			s.AttenuationAccumulation += AttenuationFunction(s.SynapseNum)
			// when s.SynapseNum is very big, AttenuationFunction(s.SynapseNum) is very small. so the more synapseNum is, the stable the link is.
			if s.AttenuationAccumulation > 1 {
				// when AttenuationAccumulation is bigger than 1, it can be turn to int so that SynapseNum can subtract it.
				s.SynapseNum -= int32(s.AttenuationAccumulation)
				if s.SynapseNum < 1 {
					// at least there is 1 Synapse, when the SynapseNum is 1, it is almost unactivatable
					s.SynapseNum = 1
				}
				s.AttenuationAccumulation = 0
			}
		}
	case config.PrefixNeureType["regulate"]:
		signalInfo := make(map[string]interface{})
		signalInfo["source_neure_id"] = thisNeure.ThisNeureId
		// 这是调节神经元的突触，不同类型的突触有不同的ActivateNextNeure方法
		if nextNeure.NeureType == config.PrefixNeureType["common"] {
			// 短时记忆就是直接让突触处于长时程增强状态，这时候的strength要加上LTPStrength
			// regulate won't activate next neure if next neure is common neure, it will regulate the linkstrength of next neure，模仿敏感化（或者叫长时程增强）
			nextSynapse := nextNeure.Synapses[s.NextNeureSynapseId]
			nextSynapse.HandleRegulationStimulate(s.LTPStrength, nextNeure)
		} else {
			signalInfo["weight"] = config.WeightThreshold + 1
			go func() {
				nextNeure.SignalChannel <- signalInfo // directly activate
			}()
		}
	case config.PrefixNeureType["inhibitory"]:
		// 抑制型神经元
	default:
		log.Panic("neure type worng!!!")
	}
	return
}

func (s *Synapse) ActivateAtOneFrequency(thisNeure *Neure, activateDuration int, activateTimes int) {
	// activateDuration是隔多少秒刺激一次，activateTimes是总共要激活多少次
	for i := 0; i < activateTimes; i++ {
		s.ActivateNextNeure(thisNeure)
		time.Sleep(time.Duration(activateDuration) * time.Second)
	}
}

func AttenuationFunction(synapseNum int32) float64 {
	// when synapseNum is very big, result is very small.
	/*
		here,AttenuationFunctionFactor means when synapseNum surpass AttenuationFunctionFactor,
		the more it is, the small result it returns
	*/
	return math.Pow(0.5, 0.05*float64((int(synapseNum)-config.AttenuationFunctionFactor)))
}

func LTPAccumulationAttenuationFunction(intervalTime float64) float64 {
	// intervalTime is the seconds between now and last activate time
	/*
		here 120 means that, when intervalTime is 120 seconds, return 0.5,
		more bigger intervalTime, more smaller, intervalTime more close to 0,
		return more close 1.
	*/
	return math.Pow(2, -(intervalTime / 120))
}
