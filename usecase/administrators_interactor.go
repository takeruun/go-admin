package usecase

import (
	"app/domain"
	"app/usecase/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdministratorsInteractor struct {
	Administrators repository.AdministratorsRepository
}

func (interactor *AdministratorsInteractor) Login(form domain.Administrator) (admin domain.Administrator, err error) {
	admin, err = interactor.Administrators.FindByEmail(form.Email)
	if err != nil {
		return domain.Administrator{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(form.Password))
	if err != nil {
		return domain.Administrator{}, err
	}

	return admin, nil
}

func (interactor *AdministratorsInteractor) Find(id uint) (admin domain.Administrator, err error) {
	admin, err = interactor.Administrators.Find(id)

	if err != nil {
		return domain.Administrator{}, err
	}

	return admin, nil
}

func (interactor *AdministratorsInteractor) FindAll() (admins []domain.Administrator, err error) {
	admins, err = interactor.Administrators.FindAll()

	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (interactor *AdministratorsInteractor) Create(form domain.Administrator) (err error) {
	ps, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	form.Password = string(ps)
	err = interactor.Administrators.Create(form)
	if err != nil {
		return err
	}

	return
}

func (interactor *AdministratorsInteractor) Update(form domain.Administrator) (err error) {
	if form.Password != "" {
		ps, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		form.Password = string(ps)

		if err != nil {
			return err
		}
	}

	err = interactor.Administrators.Update(form)
	if err != nil {
		return err
	}

	return
}
