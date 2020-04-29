package models

import (
	"github.com/napapol-dev-group/poc-golang/app/entities"
)

type RegisterUserRequest struct {
	FullName string `json:"fullName" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

// Entity convert CreateRequestJSON to Ambassador entity
func (model *RegisterUserRequest) Entity() entities.CreateUser {
	createUserEntity := entities.CreateUser{
		FullName: model.FullName,
		Address:  model.Address,
	}
	return createUserEntity
}

type RegisterUserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
