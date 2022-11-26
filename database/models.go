package database

type Neures struct {
	ID    int64  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	Neure []byte `gorm:"column:neure;type:blob;"`
}
