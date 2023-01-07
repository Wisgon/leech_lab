package database

func GetNeuresByIds(idBegin, idEnd int64) []NeureDb {
	var neures []NeureDb
	if idEnd > idBegin {
		panic("id end must bigger than id begin")
	}
	result := db.Where("id>=? AND id<=?", idBegin, idEnd).Find(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
	return neures
}

func GetNeureById(id int64) NeureDb {
	var neure NeureDb
	result := db.Where("id=?", id).First(&neure)
	if result.Error != nil {
		panic(result.Error)
	}
	return neure
}

func GetNeuresByIdArray(ids []int64) []NeureDb {
	var neures []NeureDb
	result := db.Find(&neures, ids)
	if result.Error != nil {
		panic(result.Error)
	}
	return neures
}

func GetUnlinkedNeures(amount int) []NeureDb {
	var neures []NeureDb
	result := db.Limit(amount).Where("linked = 0").Find(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
	return neures
}
