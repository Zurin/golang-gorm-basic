package repositories

import (
	"new-platform-dashboard/db"
	"new-platform-dashboard/db/entities"
)

func SaveUser(entity *entities.User) (*[]entities.User, error) {
	user := entities.User{Name: "Rino Ridlo Julianto", Email: "rinoridlojulianto@gmail.com", Password: "immortalblood", Username: "zurin"}

	err := db.DB[0].Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &[]entities.User{user}, nil
}

func FindUser(entity *entities.User) *[]entities.User {
	var users []entities.User
	db.DB[0].Find(&users)

	return &users
}
