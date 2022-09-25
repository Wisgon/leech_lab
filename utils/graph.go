package utils

import (
	"graph_robot/config"
	"graph_robot/graph_structure"
	"math/rand"
	"time"
)

func GetNewNode() *graph_structure.Node {
	rand.Seed(time.Now().UnixNano())
	rangeOfRand := config.MaxDendrites - config.MinDendrites
	node := graph_structure.Node{
		// 初始化結構體時，由config.go的最低和最高兩個設置中的隨機數生成
		MaxNumOfDendrites: rand.Intn(rangeOfRand) + int(config.MinDendrites), // Intn(n) 生成[0,n)之間的隨機整數
	}
	return &node
}

func TalkExportToString(nodes []*graph_structure.Node) (result string) {
	// 語言的output神經元轉化為string輸出
	return
}

func WordsToActivateNode(words string) (nodes []*graph_structure.Node) {
	// 文字轉化為開始的觸發神經元
	return
}
