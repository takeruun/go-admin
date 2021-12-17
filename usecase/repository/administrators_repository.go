package repository

import (
	"app/domain"
)

type AdministratorsRepository interface {
	FindAll() (admins []domain.Administrator, err error)
	FindByEmail(email string) (admin domain.Administrator, err error)
	Find(id uint) (admin domain.Administrator, err error)
	Create(form domain.Administrator) (err error)
	Update(form domain.Administrator) (err error)
	GetMyFontSize(id uint) (fontSize int, err error)
}
