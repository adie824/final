package repository

import (
	"github.com/adie824/final/final-project/model/domain"
)

type PhotoRepository interface {
	Create(photo *domain.Photo) (err error)
	GetAll() (photos []domain.Photo, err error)
	GetOne(id uint) (photo domain.Photo, err error)
	Update(photo domain.Photo) (updatedPhoto domain.Photo, err error)
	Delete(id uint) (err error)
}
