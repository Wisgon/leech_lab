package database

// ------------------------Neure database model
type NeureData struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	Neure  []byte `gorm:"column:un;type:blob;"`
	Linked bool   `gorm:"column:ed;type:boolean"`
}

func (n *NeureData) TableName() string {
	return "un"
}

func (n *NeureData) Save() int64 {
	result := db.Save(n)
	if result.Error != nil {
		panic(result.Error)
	}
	return n.ID
}

func (n *NeureData) Update() {
	result := db.Save(n)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (n *NeureData) UpdateLinked(nextId int64) {
	result := db.Model(&NeureData{}).Where("id = ?", n.ID).Update("linked", nextId)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (n *NeureData) GetNeureDataById(id int64) {
	result := db.Where("id=?", id).First(n)
	if result.Error != nil {
		panic(result.Error)
	}
}

// -------------------------NeureEntrance database model
type NeureEntrance struct {
	ID            int64  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	NeureEntrance []byte `gorm:"column:en;type:blob;"`
}

func (n *NeureEntrance) TableName() string {
	return "en"
}
