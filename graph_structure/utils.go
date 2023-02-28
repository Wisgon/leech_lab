package graph_structure

import "graph_robot/database"

func CreateOneNeure() (neure *Neure) {
	neure.CreateNeureInDB()
	return
}

func CreateNewNeures(amount int) (neures []*Neure) {
	if amount <= 0 {
		panic("amount must be lagger than 0")
	}
	emptyNeure := Neure{}
	newNeureDatasIds := database.CreateEmptyNeures(amount, emptyNeure.Struct2Byte())
	newNeureDatas := []*database.NeureData{}
	for i := 0; i < amount; i++ {
		newNeure := &Neure{ThisNeureId: newNeureDatasIds[i]}
		neures = append(neures, newNeure)
		newNeureDatas = append(newNeureDatas, &database.NeureData{
			ID:    newNeureDatasIds[i],
			Neure: newNeure.Struct2Byte(),
		})
	}
	// to update because the new NeureDatas just been created are empty, Neure byte is []byte, ThisNeureId is 0
	database.UpdateNeures(newNeureDatas)
	return
}
