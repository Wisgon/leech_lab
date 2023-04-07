package neure

func CreateOneNeure(keyPrefix string) (neure *Neure) {
	neure = &Neure{}
	neure.CreateNeureInDB(keyPrefix)
	return
}

func TurnNeureBytes2Neures(neureBytes *map[string][]byte) *map[string]*Neure {
	neures := make(map[string]*Neure)
	for k, v := range *neureBytes {
		neures[k] = &Neure{}
		neures[k].Byte2Struct(v)
	}
	return &neures
}
