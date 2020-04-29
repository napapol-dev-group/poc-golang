package user

import (
	"github.com/jinzhu/gorm"
	"github.com/napapol-dev-group/poc-golang/app/entities"
)

type repo struct {
	Conn *gorm.DB
}

// InitRepo CRUD
func InitRepo(Conn *gorm.DB) Repo {
	return &repo{Conn: Conn}
}

// UserRepo CRUD interface
type Repo interface {
	Create(user entities.CreateUser) error
}
