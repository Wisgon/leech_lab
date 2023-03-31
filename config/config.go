package config

var MinDendrites = 7    // 樹突數量的生成隨機數中的最小值
var MaxDendrites = 7777 // 樹突數量的生成隨機數中的最大值
var NeurePrefix = []string{
	"eye_entrance",
	"mouth_entrance",
}
var MaxTransactionNum = 100000 // according to experiments, this mechine can hold most 100000+ uncommit.
var FixedTransactionNum = 1000 // I think it's better every 1000 per commit.

// leech config
var DatabaseName = "leech"
var LeechSize = [3]int64{100, 100, 100}
var LeechCenterCor = [3]int64{0, 0, 0}
var BodyActions = map[string]string{
	"move_left": "move_left",
}
var DatasPath = "/home/zhilong/Documents/my_projects/graph_robot/simulate_leech/datas" // leech DatasPath
