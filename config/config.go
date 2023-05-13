package config

import (
	"graph_robot/utils"
	"time"
)

var ProjectRoot = utils.GetProjectRoot()

// global config-------------------------------------------------------------------------------------------------
var PrefixNumSplitSymbol = "@"
var PrefixNameSplitSymbol = "_"
var RefractoryDuration = 5 * time.Millisecond // 神经元的不应期

var LinkStrengthIncDuration = 5 * time.Minute // 长时程增强一次的持续时间，以分钟为单位
// var LinkStrengthIncDuration = 3 * time.Second  // debug
var BreakThroughCoefficient = float32(0.3)     // 突破系数，越大的话，与next weight越接近越容易突破
var WeightThreshold = float32(100)             // 每个神经元的激活权重都是固定的，会变化的是连接强度
var RefreshNowWeightDuration = 1 * time.Second // 神经脉冲持续时间，如果这个时间内没有能激活这个神经元，那么now weight就会重置，模仿神经元需要一段时间积累神经脉冲才能激发的特性，如果这段时间没激发，神经递质会被回收，也就是now weight被重置
var InSyncNeureMapDuration = 10 * time.Minute  // 可以在neure map里待的最长时间，超过这个时间会被存入数据库并移出map，如果经常激活的神经元太多，这个值就设置小一点
// var ActivateFrequency = 200                    //神经元激活频率，单位为次/秒
var DefaultSynapseNum = 10         //默认每个突触初始有10个分支突触与目标相连
var AttenuationFunctionFactor = 30 //衰减函数因子，当synapse大于30个的时候就开始衰减的很慢，synapse的num越大衰减得越慢
var DefaultStrength = 1.1
var LongTermMemoryLTPAThreshold = 500  // LTPA = Long-term potentiation Accumulative，即长时程增强积累数，超过这个阈值，则SynapseNum直接增加SurpassThresholdSynapseNumAdd个，按照突触的LTPStrength来看，一次增加101，那么大概五次就能超过，成为长时记忆
var SurpassThresholdSynapseNumAdd = 20 // 每次超过阈值，都会直接增加这么多个，多次超过，增加越多，就会变成长时记忆，衰减很慢了

// database setting------------------------------------------------------------------------------------------------------
var MaxTransactionNum = 100000 // according to experiments, this mechine can hold most 100000+ uncommit.
var FixedTransactionNum = 1000 // I think it's better every 1000 per commit.
var SeqBandwidth = 1024        //Setting a bandwidth too low would do more disk writes, setting it too high would result in wasted integers if Badger is closed or crashes. To avoid wasted integers, call Release before closing Badger
var PrefetchSize = 128         //By default, Badger prefetches the values of the next 100 items. You can adjust that with the IteratorOptions.PrefetchSize field. However, setting it to a value higher than GOMAXPROCS (which we recommend to be 128 or higher) shouldn’t give any additional benefits. You can also turn off the fetching of values altogether. See section below on key-only iteration.

// testing config------------------------------------------------------------------------------------------------------
var TestDataPath = ProjectRoot + "/test/datas"
var TestPrefix = "testing_neure"

// leech config------------------------------------------------------------------------------------------------------
var DatabaseName = "leech"
var LeechDatasPath = ProjectRoot + "/simulate_leech/datas" // leech DatasPath
var EachSkinPositionSurfaceNeureNum = 10                   //每个skin的area的表层神经元数量
var EachSkinPositionDeeperNeureNum = 50                    // 每个skin的area的深层神经元数量
var EachSkinPositionDeepestNeureNum = 100
var EachValuateNeureTypeNum = 1
var SignalChannelBufferSizeDefault = 2
