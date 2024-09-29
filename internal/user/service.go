package user

import (
	"Backend-Go/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Repo UserRepository
}

// Register registers a new user, ensuring unique email and username, and hashes the password
func (s *UserService) Register(user models.User) (models.User, error) {
	// Check if email is unique
	if err := s.checkUniqueEmail(user.Email); err != nil {
		return models.User{}, err
	}

	// Check if username is unique

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	// Create the user
	return s.Repo.Create(user)
}

func (s *UserService) Login(email, password string) (models.User, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	// Compare the stored hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

// checkUniqueEmail checks if the email is already taken
func (s *UserService) checkUniqueEmail(email string) error {
	_, err := s.Repo.FindByEmail(email)
	if err == nil {
		return errors.New("email already in use")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

// hashPassword hashes the password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
