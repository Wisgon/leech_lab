package database

import "strconv"

// ------------------------NeureData
func UpdateNeures(neures []NeureData) {
	result := db.Save(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
}

func GetNeuresByIds(idBegin, idEnd int64) (neures []NeureData) {
	if idEnd > idBegin {
		panic("id end must bigger than id begin")
	}
	result := db.Where("id>=? AND id<=?", idBegin, idEnd).Find(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetNeuresByIdArray(ids []int64) (neures []NeureData) {
	result := db.Find(&neures, ids)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetUnlinkedNeures(amount int) (neures []NeureData) {
	result := db.Limit(amount).Where("linked = 0").Find(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
	if len(neures) < amount {
		panic("Not enough empty Neures, need " + strconv.Itoa(amount-len(neures)) + " more")
	}
	return
}

func CreateEmptyNeures(amount int) (ids []int64) {
	var neures []NeureData
	for i := 0; i < amount; i++ {
		neures = append(neures, NeureData{})
	}
	result := db.Create(&neures)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, n := range neures {
		ids = append(ids, n.ID)
	}
	return
}
