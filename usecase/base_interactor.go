package usecase

import (
	"app/usecase/repository"
)

type BaseIntractor struct {
	Administrators repository.AdministratorsRepository
}

func (interactor *BaseIntractor) GetMyFontSize(id uint) (fontSize int, err error) {
	fontSize, err = interactor.Administrators.GetMyFontSize(id)

	if err != nil {
		return 0, err
	}

	return fontSize, nil
}

func ReturnIntPointer(value int) *int {
	return &value
}
