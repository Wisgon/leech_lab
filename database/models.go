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

func (n *NeureData) Create() int64 {
	result := db.Create(n)
	if result.Error != nil {
		panic(result.Error)
	}
	return n.ID
}

func (n *NeureData) Save() int64 {
	result := db.Save(n)
	if result.Error != nil {
		panic(result.Error)
	}
	return n.ID
}

func (n *NeureData) GetNeureDataById(id int64) {
	result := db.Where("id=?", id).First(n)
	if result.Error != nil {
		panic(result.Error)
	}
}
