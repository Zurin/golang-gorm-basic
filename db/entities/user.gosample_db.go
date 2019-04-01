package entities

type User struct {
	Id       uint `gorm:"AUTO_INCREMENT"`
	Name     string
	Email    string
	Username string
	Password string
}
