package human

// import (
// 	"graph_robot/neure"
// )

// //小知识
// // 眼睛的视觉皮层的神经元要设置一个固定像素，每个像素点连接一个神经元，将像素值给转换成神经冲动

// type HumanBrain struct {
// 	EyesEntrance     neure.NeureEntrance // 視覺神經的觸發神經元
// 	EarEntrance      neure.NeureEntrance // 聽覺神經的觸發神經元
// 	LanguageEntrance neure.NeureEntrance // language input entrance
// 	SensoryEntrance  neure.NeureEntrance // sensory entrance is use for get the signal from sensory organ
// 	ValueArea        neure.NeureEntrance // 價值判斷區，負責給與反饋，反饋大小取決於給與的電流強弱（分數的正負，負數越多越不好，到某種程度就是恐懼感,正數越大越好），是一組預先定義好的神經元,如果某個動作是正確的，則要把這個動作的神經元組連到這個區域中鞏固下來,以後想到這個動作就會得到這個value反饋
// 	SenceArea        neure.NeureEntrance // 情景网络
// 	MemoryArea       neure.NeureEntrance // 记忆网络
// }

// func (b *HumanBrain) InitBrain() {
// }

// func (b *HumanBrain) Think() {
// 	// 啟動一個永不停歇的攜程，作為思考的一個神經衝動
// 	// 靈感是隨機觸發think相關的神經元，以觸發神經元為起點，終點為思考結果output
// 	// 可以與價值區域相關聯，如果想到之前的某段情景記憶，可以觸發價值判斷區域，就感受到了這個情景的好壞的感受
// 	// 大概率會去執行那些好的感受，小概率執行中性的，不執行不好的感受，極力避開非常不好的感受
// }

// func (b *HumanBrain) TalkExportToString() (result string) {
// 	// 語言的output神經元轉化為string輸出
// 	return
// }

// func (b *HumanBrain) WordsToActivateNeure(words string) {
// 	// 文字轉化為開始的觸發神經元
// }

// func (b *HumanBrain) Output2ValueNetwork(resultNetwork neure.NeureEntrance, value int) {
// 	// 这是把神经元的output和value网络连接起来，value是这个result返回的结果，可以是人工返回的，也可以是环境返回的
// 	// 这个方法将结果和value网络连接起来，每次触发出这个result，都会连接到这个value，如果相同的事情，环境返回了不同的value，则会加强或者削弱value值
// 	// var newValueNeureEntrance neure.NeureEntrance
// }

// func (b *HumanBrain) Output2SceneNetwork(resultNetwork neure.NeureEntrance) {
// 	// 这是把output链接到场景神经元的方法，也是记忆单词的必要步骤，把单词链接到学单词时的场景
// 	// 目前还不清楚场景是否就是视觉记忆，有待思考：todo:
// 	// var newSceneNeureEntrance neure.NeureEntrance
// }
