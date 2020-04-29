package user

import (
	"errors"
	"github.com/napapol-dev-group/poc-golang/app/entities"
)

func (u useCase) RegisterUser(createUserEntity entities.CreateUser) (*entities.CreateUser, error) {
	if createUserEntity.Address != "thai" {
		return nil, errors.New("address is not support")
	}
	err := u.UserRepo.Create(createUserEntity)
	return &createUserEntity, err
}
