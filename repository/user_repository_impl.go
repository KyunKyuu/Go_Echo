package repository

import (
	"rest_gorm/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) SaveUser(user domain.User) (domain.User, error) {

	err := repo.db.Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUser(Id int) (domain.User, error) {
	var userData domain.User

	err := repo.db.First(&userData, "user_id = ?", Id).Error

	if err != nil {
		return domain.User{}, err
	}

	return userData, nil
}

func (repo *UserRepositoryImpl) GetUserDeleted(Id int) (domain.User, error) {
	var userData domain.User

	err := repo.db.Unscoped().Where("user_id = ?", Id).Find(&userData).Error

	if err != nil {
		return domain.User{}, err
	}

	return userData, nil

}

func (repo *UserRepositoryImpl) GetUsers() ([]domain.User, error) {
	var users []domain.User

	err := repo.db.Find(&users).Error

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (repo *UserRepositoryImpl) UpdateUser(user domain.User) (domain.User, error) {
	err := repo.db.Model(domain.User{}).Where("user_id = ?", user.UserID).Updates(user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) DeleteUser(Id int) error {

	var userData domain.User

	err := repo.db.Delete(&userData, Id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) FindUserByEmail(email string) (*domain.User, error) {
	user := new(domain.User)

	if err := repo.db.Where("email=?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
