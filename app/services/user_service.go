package services

import (
	"demo/app/models"
	"demo/app/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (userService *UserService) GetAllUsers() ([]models.User, error) {
	return userService.UserRepository.GetAllUsers()
}
