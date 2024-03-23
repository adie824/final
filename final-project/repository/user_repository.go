package repository

import (
	"github.com/adie824/final/final-project/model/domain"
)

type UserRepository interface {
	Register(user *domain.User) (err error)
	Login(user *domain.User) (err error)
	Update(user *domain.User) (updatedUser domain.User, err error)
	Delete(id uint) (err error)
}
