package human

import (
	"fmt"
	"graph_robot/neure"
)

type Mouth struct {
}

func (m *Mouth) Convert(outputNeures []neure.Neure) string {
	return ""
}

func (m *Mouth) Say() {
	fmt.Println("")
}

type Body struct {
}

// todo:temporally output a string like "touch", "hit" and so on, finally will connect to real action.
func (b *Body) Convert(outputNeures []neure.Neure) string {
	return ""
}

func (b *Body) Action() {

}
