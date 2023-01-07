package database

type NeureDb struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	Neure  []byte `gorm:"column:neure;type:blob;"`
	Linked bool   `gorm:"column:linked;type:boolean"`
}

func (n *NeureDb) TableName() string {
	return "neures"
}
