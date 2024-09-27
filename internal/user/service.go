package user

import "Backend-Go/internal/models"

type UserService struct {
	Repo UserRepository
}

func (s *UserService) Register(user models.User) (models.User, error) {
	// Add validation and hashing password logic here
	return s.Repo.Create(user)
}
