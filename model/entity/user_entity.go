package entity

import (
	"rest_gorm/model/domain"
)

type UserEntity struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func ToUserEntity(user_id int, name string, email string) UserEntity {
	return UserEntity{
		UserID: user_id,
		Name:   name,
		Email:  email,
	}
}

func ToUserListEntity(users []domain.User) []UserEntity {
	userData := []UserEntity{}

	for _, user := range users {
		userData = append(userData, ToUserEntity(user.UserID, user.Name, user.Email))
	}

	return userData
}
