package neure

func CreateOneNeure(keyPrefix string) (neure *Neure) {
	neure = &Neure{}
	neure.CreateNeureInDB(keyPrefix)
	return
}
