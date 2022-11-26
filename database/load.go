package database

func GetNeuresByIds(idBegin, idEnd int64) []Neures {
	var neures []Neures
	if idEnd > idBegin {
		panic("id end must bigger than id begin")
	}
	result := db.Where("id>=? AND id<=?", idBegin, idEnd).Find(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
	return neures
}

func GetNeureById(id int64) Neures {
	var neure Neures
	result := db.Where("id=?", id).First(&neure)
	if result.Error != nil {
		panic(result.Error)
	}
	return neure
}

func GetNeuresByIdArray(ids []int64) []Neures {
	var neures []Neures
	result := db.Find(&neures, ids)
	if result.Error != nil {
		panic(result.Error)
	}
	return neures
}
