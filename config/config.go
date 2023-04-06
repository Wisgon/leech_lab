package config

import "graph_robot/utils"

var ProjectRoot = utils.GetProjectRoot()

// global config-----------------------------------------------
var MinDendrites = 7    // 樹突數量的生成隨機數中的最小值
var MaxDendrites = 7777 // 樹突數量的生成隨機數中的最大值MaxDendrites
// 神经元的Prefix命名原则是：功能名+entrance（除了这个，还可以是normal，output或其他）
// 不一定要这个命名原则，反正只要描述清楚神经元的功能或位置就行
var NeurePrefix = []string{
	"testing_neure",

	// -----eyes--------
	"eye_entrance",
	"eye_normal",
	"eye_output",

	// ------mouth------
	"mouth_entrance",
	"mouth_normal",
	"moutn_output",

	// -------nose----------
	"nose_entrance",
	"nose_normal",
	"nose_output",

	// --------sense------
	"sense_entrance",
	"sense_normal",
	"sense_output",

	// ----------sense memory--------
	"senseMemory_entrance",
	"senseMemory_normal",
	"senseMemory_output",

	// --------scene memory------
	"sceneMemory_entrance",
	"sceneMemory_normal",
	"sceneMemory_output",

	// ----------smell memory-------
	"smellMemory_entrance",
	"smellMemory_normal",
	"smellMemory_output",

	// --------short term memory-----
	"shortTermMemory_entrance",
	"shortTermMemory_normal",
	"shortTermMemory_output",

	// ----------valuate----------
	"valuate_entrance",
	"valuate_normal",
	"valuate_output",

	// -----------skin-----------
	"skin_entrance",
	"skin_normal",
	"skin_output",

	// ------self consciousness----------
	"selfConsciousness_entrance",
}
var PrefixNumSplitSymbol = "@"

// database setting----------------------------------------------------
var MaxTransactionNum = 100000 // according to experiments, this mechine can hold most 100000+ uncommit.
var FixedTransactionNum = 1000 // I think it's better every 1000 per commit.
var SeqBandwidth = 1024        //Setting a bandwidth too low would do more disk writes, setting it too high would result in wasted integers if Badger is closed or crashes. To avoid wasted integers, call Release before closing Badger
var PrefetchSize = 128         //By default, Badger prefetches the values of the next 100 items. You can adjust that with the IteratorOptions.PrefetchSize field. However, setting it to a value higher than GOMAXPROCS (which we recommend to be 128 or higher) shouldn’t give any additional benefits. You can also turn off the fetching of values altogether. See section below on key-only iteration.

// testing config----------------------------------------------
var TestDataPath = ProjectRoot + "/test/datas"

// leech config------------------------------------------------
var DatabaseName = "leech"
var LeechSize = [3]int64{100, 100, 100}
var LeechCenterCor = [3]int64{0, 50, 0}
var BodyActions = map[string]string{
	"move_left": "move_left",
}
var LeechDatasPath = ProjectRoot + "/simulate_leech/datas" // leech DatasPath
var SkinNeurePosition = []string{"left_front_up", "left_front_down"}
