package brain

import gs "graph_robot/graph_structure"

type Brain struct {
	Eyes []gs.NetWork // 視覺神經的觸發神經元
	// Ear  []gs.NetWork // 聽覺神經的觸發神經元
	Mouth      []gs.NetWork // word output network
	ValueArea  []gs.NetWork // 價值判斷區，負責給與反饋，反饋大小取決於給與的電流強弱（分數的正負，負數越多越不好，到某種程度就是恐懼感,正數越大越好），是一組預先定義好的神經元,如果某個動作是正確的，則要把這個動作的神經元組連到這個區域中鞏固下來,以後想到這個動作就會得到這個value反饋
	SenceArea  []gs.NetWork // 情景网络
	MemoryArea []gs.NetWork // 记忆网络
}

func (b *Brain) TalkExportToString() (result string) {
	// 語言的output神經元轉化為string輸出
	return
}

func (b *Brain) WordsToActivateNeure(words string) {
	// 文字轉化為開始的觸發神經元
}

func (b *Brain) Think() {
	// 啟動一個永不停歇的攜程，作為思考的一個神經衝動
	// 靈感是隨機觸發think相關的神經元，以觸發神經元為起點，終點為思考結果output
	// 可以與價值區域相關聯，如果想到之前的某段情景記憶，可以觸發價值判斷區域，就感受到了這個情景的好壞的感受
	// 大概率會去執行那些好的感受，小概率執行中性的，不執行不好的感受，極力避開非常不好的感受
}
