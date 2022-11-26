package graph_structure

type BrainGraph struct {
	EyeEntries  []int // 視覺神經的觸發神經元
	TalkExport  []int // 說話的output
	EarEntries  []int // 聽覺神經的觸發神經元
	WordEntries []int // 閱讀文字的觸發神經元
}

func (b *BrainGraph) Think() {
	// 啟動一個永不停歇的攜程，作為思考的一個神經衝動
	// 靈感是隨機觸發think相關的神經元，以觸發神經元為起點，終點為思考結果output
}

func (b *BrainGraph) Thalamus() {
	// 丘腦區，丘腦是負責匯集身體各部分的感覺器官傳過來的神經衝動，然後傳給大腦皮層處理
}
