package services

import "github.com/ipan97/mumu-store/app/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}
