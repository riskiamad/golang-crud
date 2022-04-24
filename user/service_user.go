package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(userInput UserInput) (User, error)
	Update(ID int, userInput UserInput) (User, error)
	Delete(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("the user with that is not found")
	}

	return user, nil
}

func (s *service) Create(userInput UserInput) (User, error) {
	user := User{
		Username:  userInput.Username,
		Email:     userInput.Email,
		CreatedAt: time.Now(),
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Create(user)
	if err != nil {
		return user, err
	}

	return newUser, err
}

func (s *service) Update(ID int, userInput UserInput) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("the user with that ID is not found")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Username = userInput.Username
	user.Email = userInput.Email
	user.Password = string(passwordHash)
	user.UpdatedAt = time.Now()

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, err
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("the user with that ID is not found")
	}

	deleteUser, err := s.repository.Delete(user)
	if err != nil {
		return deleteUser, err
	}

	return deleteUser, err
}
