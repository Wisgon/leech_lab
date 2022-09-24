package graph_structure

// 小知識：
// 1. 神經元神經衝動傳導方向時樹突傳向軸突然後傳到下一個神經元的樹突
// 2. 一個神經元的軸突只能連接一個神經元的樹突，但是樹突可以連接多個神經元，將這些神經元的信號通過軸突傳給下一個神經元

type Node struct {
	NextNode          *Node // 軸突連接的下一個神經元
	Weight            int   // 與nextNode的連接權重
	DendritesLinkNum  int   // 樹突的數量
	MaxNumOfDendrites int   // 樹突的最大數量，不能超過這個數量的連接
}

func (n *Node) DecreaseDendritesNum() {
	// 上一個神經元的軸突與該神經元的樹突斷開連接時要減一
	if n.DendritesLinkNum > 0 {
		n.DendritesLinkNum -= 1
	}
}

func (n *Node) IncreaseDendritesNum() {
	// 神經元的樹突與其他神經元的軸突連接時要加1
	if n.DendritesLinkNum < n.MaxNumOfDendrites {
		n.DendritesLinkNum += 1
	}
}
