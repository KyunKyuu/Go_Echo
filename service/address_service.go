package service

import (
	"rest_gorm/helper"
	"rest_gorm/model/entity"
	"rest_gorm/model/web"
)

type AddressService interface {
	Create(token string, req web.AddressServiceRequest) (helper.ResponseToJson, error)
	Update(token string, id int, req web.AddressUpdateRequest) (helper.ResponseToJson, error)
	Delete(id int) error
	GetAddress(token string) (entity.AddressEntity, error)
	GetAllAddress() ([]entity.AddressEntity, error)
	GetDetail(id int) (entity.DetailAddress, error)
}
