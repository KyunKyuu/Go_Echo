package service

import (
	"rest_gorm/model/entity"
	"rest_gorm/model/web"
)

type UserService interface {
	SaveUser(request web.UserServiceRequest) (map[string]interface{}, error)
	GetUser(userId int) (entity.UserEntity, error)
	GetUserDeleted(userId int) (entity.UserEntity, error)
	GetUserList() ([]entity.UserEntity, error)
	UpdateUser(request web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	DeleteUser(userId int) error
	LoginUser(email string, password string) (map[string]interface{}, error)
}
