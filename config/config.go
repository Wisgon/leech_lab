package config

var MinDendrites = 7    // 樹突數量的生成隨機數中的最小值
var MaxDendrites = 7777 // 樹突數量的生成隨機數中的最大值
var EntranceTypes = map[string]string{
	"eyes":  "eyes",
	"mouth": "mouth",
}

// leech config
var DatabaseName = "leech"
var LeechSize = [3]int64{100, 100, 100}
var LeechCenterCor = [3]int64{0, 0, 0}
var BodyActions = map[string]string{
	"move_left": "move_left",
}
