package user

import (
	"Backend-Go/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
