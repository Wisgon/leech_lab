package graph_structure

import (
	"encoding/json"
	"graph_robot/utils"
	"os"
)

type NeureEntrance struct {
	EntranceType string  // like "eyes", "ears", "mouth" and so on
	NeuresIds    []int64 // Start neures ids of the network
}

func (n *NeureEntrance) Save2File() {
	err := os.WriteFile(n.GetEntranceFilePath(), n.Struct2Byte(), 0644)
	if err != nil {
		panic(err)
	}
}

func (n *NeureEntrance) LoadFromFile() {
	data, err := os.ReadFile(n.GetEntranceFilePath())
	if err != nil {
		panic(err)
	}
	n.Byte2Struct(data)
}

func (n *NeureEntrance) GetEntranceFilePath() string {
	rootPath := utils.GetProjectRoot()
	filePath := rootPath + "/entrances/" + n.EntranceType
	return filePath
}

func (n *NeureEntrance) Struct2Byte() []byte {
	nb, err := json.Marshal(n)
	if err != nil {
		panic("json marshal error: " + err.Error())
	}
	return nb
}

func (n *NeureEntrance) Byte2Struct(neureByte []byte) {
	err := json.Unmarshal(neureByte, n)
	if err != nil {
		panic("json unmarshal error: " + err.Error())
	}
}
