package database

import (
	"app/domain"
)

type ProductsRepository struct {
	DB DB
}

func (repo *ProductsRepository) FindAll(offset int, limit int, form domain.Product) (products []domain.Product, err error) {
	db := repo.DB.Connect()

	result := db.Where(&form).Preload("Category").Offset(offset).Limit(limit).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *ProductsRepository) FindById(id int) (product domain.Product, err error) {
	db := repo.DB.Connect()

	result := db.Preload("Category").Find(&product, id)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}

	return product, nil
}

func (repo *ProductsRepository) Create(form domain.Product) (err error) {
	db := repo.DB.Connect()

	result := db.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return
}

func (repo *ProductsRepository) Update(form domain.Product) (err error) {
	db := repo.DB.Connect()

	result := db.Updates(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ProductsRepository) Search(form domain.Product) (prodcuts []domain.Product, err error) {
	db := repo.DB.Connect()

	result := db.Where(&form).Find(&prodcuts)
	if result.Error != nil {
		return nil, result.Error
	}

	return prodcuts, nil
}

func (repo *ProductsRepository) TotalCount(form domain.Product) (count int64, err error) {
	db := repo.DB.Connect()

	result := db.Model(domain.Product{}).Where(&form).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return
}
