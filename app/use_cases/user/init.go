package user

import (
	"github.com/napapol-dev-group/poc-golang/app/entities"
	repository "github.com/napapol-dev-group/poc-golang/app/repositories/user"
)

type useCase struct {
	UserRepo repository.Repo
}

// InitUseCase init Operate Area use case
func InitUseCase(userRepository repository.Repo) UseCase {
	return &useCase{
		UserRepo: userRepository,
	}
}

// UseCase Operate Area use case interface
type UseCase interface {
	RegisterUser(createUser entities.CreateUser) (*entities.CreateUser, error)
}
