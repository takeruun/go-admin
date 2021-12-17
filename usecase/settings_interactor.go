package usecase

import (
	"app/domain"
	"app/usecase/repository"
)

type SettingsInteractor struct {
	Settings       repository.SettingsRepository
	Administrators repository.AdministratorsRepository
}

func (interactor *SettingsInteractor) GetMyFontSize(id uint) (fontSize int, err error) {
	admin, err := interactor.Administrators.Find(id)
	if err != nil {
		return 0, err
	}

	return admin.FontSize, nil
}

func (interactor *SettingsInteractor) UpdateFontSize(form domain.Administrator) (fontSize int, err error) {
	err = interactor.Administrators.Update(form)
	if err != nil {
		return 0, err
	}

	admin, err := interactor.Administrators.Find(form.ID)
	if err != nil {
		return 0, err
	}

	return admin.FontSize, nil
}
