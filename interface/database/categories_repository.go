package database

import "app/domain"

type CategoriesRepository struct {
	DB DB
}

func (repo *CategoriesRepository) FindAll(offset int, limit int) (categories []domain.Category, err error) {
	db := repo.DB.Connect()

	result := db.Offset(offset).Limit(limit).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (repo *CategoriesRepository) Find(id int) (category domain.Category, err error) {
	db := repo.DB.Connect()

	result := db.First(&category, id)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}

	return category, nil
}

func (repo *CategoriesRepository) Create(form domain.Category) (err error) {
	db := repo.DB.Connect()

	result := db.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return
}

func (repo *CategoriesRepository) Update(form domain.Category) (err error) {
	db := repo.DB.Connect()

	result := db.Updates(&form)
	if result.Error != nil {
		return result.Error
	}

	return
}

func (repo *CategoriesRepository) TotalCount() (count int64, err error) {
	db := repo.DB.Connect()

	result := db.Model(domain.Category{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return
}
