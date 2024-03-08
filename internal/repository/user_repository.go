package repository

import (
	"codecreeo/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) Update(user *model.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(user *model.User) error {
	return ur.db.Delete(user).Error
}

func (ur *UserRepository) GetById(id uint) (*model.User, error) {
	var user model.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
