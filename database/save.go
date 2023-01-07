package database

func UpdateNeures(neures []NeureDb) {
	result := db.Save(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
}

func SaveNeure(neure NeureDb) int64 {
	result := db.Save(&neure)
	if result.Error != nil {
		panic(result.Error)
	}
	return neure.ID
}

func UpdateLinked(neureId int64) {
	result := db.Model(&NeureDb{}).Where("id = ?", neureId).Update("linked", 1)
	if result.Error != nil {
		panic(result.Error)
	}
}
