package graph_structure

type BrainGraph struct {
	EyeEntries  []*Node // 視覺神經的觸發神經元
	TalkExport  []*Node // 說話的output
	EarEntries  []*Node // 聽覺神經的觸發神經元
	ReadEntries []*Node // 閱讀文字的觸發神經元
}

func (b *BrainGraph) Think() {
	// 啟動一個永不停歇的攜程，作為思考的一個神經衝動
	// 靈感是隨機觸發think相關的神經元，以觸發神經元為起點，終點為思考結果output
}
