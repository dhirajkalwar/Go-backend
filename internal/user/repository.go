package user

import (
	"Backend-Go/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	result := r.DB.Create(&user)
	return user, result.Error
}
