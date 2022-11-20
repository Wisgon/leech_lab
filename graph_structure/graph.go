package graph_structure

import (
	"graph_robot/config"
	"graph_robot/utils"
	"math/rand"
	"sync"
	"time"
)

var Brain sync.Map

type BrainGraph struct {
	EyeEntries  []*Neure // 視覺神經的觸發神經元
	TalkExport  []*Neure // 說話的output
	EarEntries  []*Neure // 聽覺神經的觸發神經元
	WordEntries []*Neure // 閱讀文字的觸發神經元
}

func (b *BrainGraph) GetNewNeure(NeureType bool) string {
	nowNano := time.Now().UnixNano()
	rand.Seed(nowNano)
	rangeOfRand := config.MaxDendrites - config.MinDendrites
	dendritesLinkNum := rand.Intn(rangeOfRand) + int(config.MinDendrites) // Intn(n) 生成[0,n)之間的隨機整數
	hashID := utils.GetUniqueId(nowNano)
	Neure := Neure{
		// 初始化結構體時，由config.go的最低和最高兩個設置中的隨機數生成
		DendritesLinkNum: dendritesLinkNum,
		NeureType:        NeureType,
		HashId:           hashID,
	}
	Brain.Store(hashID, Neure)
	return hashID
}

func (b *BrainGraph) Think() {
	// 啟動一個永不停歇的攜程，作為思考的一個神經衝動
	// 靈感是隨機觸發think相關的神經元，以觸發神經元為起點，終點為思考結果output
}

func (b *BrainGraph) Thalamus() {
	// 丘腦區，丘腦是負責匯集身體各部分的感覺器官傳過來的神經衝動，然後傳給大腦皮層處理
}
