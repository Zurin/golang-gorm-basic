package entities

type Example struct {
	Id      uint   `gorm:"AUTO_INCREMENT"`
	Message string `gorm:"type:varchar(100)"`
}
