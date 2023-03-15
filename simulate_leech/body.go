package leech

import "graph_robot/neure"

type LeechBody struct {
}

func (b *LeechBody) Convert(outputNeures []neure.Neure) string {
	return ""
}

func (b *LeechBody) Action() {
	// todo:temporally output a string like "touch", "hit" and so on, finally will connect to real action.
	// order := b.Convert()
}
