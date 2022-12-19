package graph_structure

type BrainGraph struct {
	// EyeEntries  []NetWork // 視覺神經的觸發神經元
	TalkExport []NetWork // 說話的output
	// EarEntries  []NetWork // 聽覺神經的觸發神經元
	WordEntries []NetWork // 閱讀文字的觸發神經元
	ValueArea   []NetWork // 價值判斷區，負責給與反饋，反饋大小取決於給與的電流強弱（分數的正負，負數越多越不好，到某種程度就是恐懼感,正數越大越好），是一組預先定義好的神經元,如果某個動作是正確的，則要把這個動作的神經元組連到這個區域中鞏固下來,以後想到這個動作就會得到這個value反饋
	SenceArea   []NetWork // 情景网络
	MemoryArea  []NetWork // 记忆网络
}

func (b *BrainGraph) TalkExportToString() (result string) {
	// 語言的output神經元轉化為string輸出
	return
}

func (b *BrainGraph) WordsToActivateNeure(words string) {
	// 文字轉化為開始的觸發神經元
	return
}
