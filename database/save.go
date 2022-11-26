package database

func UpdateNeures(neures []Neures) {
	result := db.Save(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
}

func SaveNeure(neure Neures) int64 {
	result := db.Save(&neure)
	if result.Error != nil {
		panic(result.Error)
	}
	return neure.ID
}
