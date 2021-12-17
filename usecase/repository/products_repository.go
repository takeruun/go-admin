package repository

import "app/domain"

type ProductsRepository interface {
	FindAll(offset int, limit int, form domain.Product) (products []domain.Product, err error)
	FindById(id int) (product domain.Product, err error)
	Create(form domain.Product) (err error)
	Update(form domain.Product) (err error)
	Search(form domain.Product) (prodcuts []domain.Product, err error)
	TotalCount(form domain.Product) (count int64, err error)
}
