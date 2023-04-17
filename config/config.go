package config

import "graph_robot/utils"

var ProjectRoot = utils.GetProjectRoot()

// global config-----------------------------------------------
// 神经元的Prefix命名原则是：功能名+entrance（除了这个，还可以是normal，output或其他）
// 不一定要这个命名原则，反正只要描述清楚神经元的功能或位置就行
var PrefixFirst = []string{
	"eye",
	"mouth",
	"nose",
	"sense",
	"scene", //情景处理
	"valuate",
	"skin",
	// "selfConsciousness",
	"movement",
	"reading",
	"talking",
}

var PrefixSecond = []string{
	"normal",
	"regulate", //调控神经元，是用来增强目标神经元的突触连接的
}
var PrefixSkinThird = []string{
	"normalTemperature",
	"hotTemperature",
	"coldTemperature",
	"extremelyHotTemp",
	"extremelyColdTemp",
}
var PrefixNumSplitSymbol = "@"
var PrefixNameSplitSymbol = "_"

// database setting----------------------------------------------------
var MaxTransactionNum = 100000 // according to experiments, this mechine can hold most 100000+ uncommit.
var FixedTransactionNum = 1000 // I think it's better every 1000 per commit.
var SeqBandwidth = 1024        //Setting a bandwidth too low would do more disk writes, setting it too high would result in wasted integers if Badger is closed or crashes. To avoid wasted integers, call Release before closing Badger
var PrefetchSize = 128         //By default, Badger prefetches the values of the next 100 items. You can adjust that with the IteratorOptions.PrefetchSize field. However, setting it to a value higher than GOMAXPROCS (which we recommend to be 128 or higher) shouldn’t give any additional benefits. You can also turn off the fetching of values altogether. See section below on key-only iteration.

// testing config----------------------------------------------
var TestDataPath = ProjectRoot + "/test/datas"
var TestPrefix = "testing_neure"

// leech config------------------------------------------------
var DatabaseName = "leech"
var LeechSize = [3]int64{100, 100, 100}
var LeechCenterCor = [3]int64{0, 50, 0}
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
var SkinNeurePosition = []string{
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
var EachSkinPositionNeureNum = 10
