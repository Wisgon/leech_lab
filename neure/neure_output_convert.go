package neure

type Converter interface {
	Convert(outputNeures []Neure) string
}
