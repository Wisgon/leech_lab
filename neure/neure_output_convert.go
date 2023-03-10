package neure

type Mouth interface {
	Say(outputNeures []Neure) string
}

type Body interface {
	// todo:temporally output a string like "touch", "hit" and so on, finally will connect to real action.
	Action(outputNeures []Neure) string
}
