package config

import (
	"graph_robot/utils"
	"time"
)

var ProjectRoot = utils.GetProjectRoot()

// global config-------------------------------------------------------------------------------------------------
// 神经元的Prefix命名原则是：功能名+entrance（除了这个，还可以是normal，output或其他）
// 不一定要这个命名原则，反正只要描述清楚神经元的功能或位置就行
var PrefixArea = []string{
	"eye",
	"mouth",
	"nose",
	"sense",
	"valuate",
	"skin",
	"consciousness",
	"movement",
	"reading",
	"talking",
	"hippocampus", // 海马体
}
var PrefixNeureType = []string{
	"inborn",   // 天生的
	"acquired", // 后天获得的
}
var PrefixSkinAndSenseType = []string{
	"normalTemperature",
	"hotTemperature",
	"coldTemperature",
	"extremelyHotTemp",
	"extremelyColdTemp",
	"normalPress",
	"biggerPress",
	"extremelyPress",
}
var PrefixSenseType = []string{
	"touchType",
	"painfulType", // painfulType had more higher weight
}
var PrefixNumSplitSymbol = "@"
var PrefixNameSplitSymbol = "_"
var RefractoryPeriod = 5 * time.Millisecond // 神经元的不应期

var LinkStrengthInc = 3.2                  // 长时程增强一次的强度
var LinkStrengthIncTime = 60 * time.Minute // 长时程增强一次的时间，以分钟为单位
var BreakThroughCoefficient = float32(0.3) // 突破系数，越大的话，与next weight越接近越容易突破
var Weight = float32(100)                  // 每个神经元的激活权重都是固定的，会变化的是连接强度

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
var Movements = []string{
	"move_left_front_up",
	"move_left_front_down",
	"move_left_middle_up",
	"move_left_middle_down",
	"move_left_back_up",
	"move_left_back_down",
	"move_right_front_up",
	"move_right_front_down",
	"move_right_middle_up",
	"move_right_middle_down",
	"move_right_back_up",
	"move_right_back_down",
}
var LeechDatasPath = ProjectRoot + "/simulate_leech/datas" // leech DatasPath
var SkinAndSenseNeurePosition = []string{
	"left_front_up",
	"left_front_down",
	"left_middle_up",
	"left_middle_down",
	"left_back_up",
	"left_back_down",
	"right_front_up",
	"right_front_down",
	"right_middle_up",
	"right_middle_down",
	"right_back_up",
	"right_back_down",
}
var EachSkinPositionSurfaceNeureNum = 10 //每个skin的area的表层神经元数量
var EachSkinPositionDeepNeureNum = 100   // 每个skin的area的深层神经元数量
