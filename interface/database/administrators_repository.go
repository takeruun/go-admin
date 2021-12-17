package database

import (
	"app/domain"
)

type AdministratorsRepository struct {
	DB DB
}

func (repo *AdministratorsRepository) FindAll() (admins []domain.Administrator, err error) {
	db := repo.DB.Connect()

	result := db.Find(&admins)
	if result.Error != nil {
		return nil, result.Error
	}

	return admins, nil
}

func (repo *AdministratorsRepository) FindByEmail(email string) (admin domain.Administrator, err error) {
	db := repo.DB.Connect()
	result := db.First(&admin, "email = ?", email)
	if result.Error != nil {
		return domain.Administrator{}, result.Error
	}

	return admin, nil
}

func (repo *AdministratorsRepository) Find(id uint) (admin domain.Administrator, err error) {
	db := repo.DB.Connect()

	result := db.First(&admin, id)
	if result.Error != nil {
		return domain.Administrator{}, result.Error
	}

	return admin, nil
}

func (repo *AdministratorsRepository) GetMyFontSize(id uint) (fontSize int, err error) {
	db := repo.DB.Connect()
	var admin domain.Administrator

	result := db.First(&admin, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return admin.FontSize, nil
}

func (repo *AdministratorsRepository) Create(form domain.Administrator) (err error) {
	db := repo.DB.Connect()

	result := db.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *AdministratorsRepository) Update(form domain.Administrator) (err error) {
	db := repo.DB.Connect()

	result := db.Updates(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
