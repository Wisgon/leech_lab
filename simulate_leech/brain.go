package leech

type LeechBrain struct {
}

func (lb *LeechBrain) InitBrain() {
	go lb.initSenseArea()
	go lb.initSceneMemoryArea()
	go lb.initSmellMemoryArea()
	go lb.initSenseMemoryArea()
	go lb.initShortTermMemoryArea()
	go lb.initValueArea()
	go lb.initSelfConsciousness()
}

func (lb *LeechBrain) initSenseArea() {

}

func (lb *LeechBrain) initSenseMemoryArea() {

}

func (lb *LeechBrain) initSmellMemoryArea() {

}

func (lb *LeechBrain) initSceneMemoryArea() {
}

func (lb *LeechBrain) initShortTermMemoryArea() {

}

func (lb *LeechBrain) initValueArea() {

}

func (lb *LeechBrain) initSelfConsciousness() {

}

func (lb *LeechBrain) Sense2Action() (bodyAction string) {
	return ""
}
