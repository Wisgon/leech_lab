package neure

func CreateOneNeure[T ST](keyPrefix string) (neure *Neure[T]) {
	neure = &Neure[T]{}
	neure.CreateNeureInDB(keyPrefix)
	return
}

func TurnNeureBytes2Neures[T ST](neureBytes *map[string][]byte) *map[string]*Neure[T] {
	neures := make(map[string]*Neure[T])
	for k, v := range *neureBytes {
		neures[k] = &Neure[T]{}
		neures[k].Byte2Struct(v)
	}
	return &neures
}

func RemoveUniqueValueFromSynapse[T ST](value string, s []T) []T {
	for i, v := range s {
		if v.GetNextId() == value {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}
