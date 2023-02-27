package graph_structure

import "graph_robot/database"

func CreateNewNeures(amount int) (neures []Neure) {
	if amount <= 0 {
		panic("amount must be lagger than 0")
	}
	newNeureDatasIds := database.CreateEmptyNeures(amount)
	neureDatas := database.GetNeuresByIds(newNeureDatasIds[0], newNeureDatasIds[len(newNeureDatasIds)-1])
	for i := 0; i < amount; i++ {
		neures = append(neures, Neure{DatabaseModel: neureDatas[i]})
	}
	return
}
