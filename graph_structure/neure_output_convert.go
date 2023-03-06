package graph_structure

type Mouth struct{}

func (m *Mouth) Say(outputNeures []Neure) string {
	return ""
}

type Body struct{}

func (b *Body) Action(outputNeures []Neure) string {
	// todo:temporally output a string like "touch", "hit" and so on, finally will connect to real action.
	return ""
}
