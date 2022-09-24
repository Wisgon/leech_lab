package utils

import (
	"graph_robot/graph_structure"
	"time"

	"math/rand"
)

func GetNewNode() *graph_structure.Node {
	rand.Seed(time.Now().UnixNano())
	rangeOfRand := graph_structure.MaxDendrites - graph_structure.MinDendrites
	node := graph_structure.Node{
		// 初始化結構體時，由config.go的最低和最高兩個設置中的隨機數生成
		MaxNumOfDendrites: rand.Intn(rangeOfRand) + int(graph_structure.MinDendrites), // Intn(n) 生成[0,n)之間的隨機整數
	}
	return &node
}

func TalkExportToString(nodes []*graph_structure.Node) (result string) {
	// 語言的output神經元轉化為string輸出
	return
}
