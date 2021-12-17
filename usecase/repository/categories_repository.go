package repository

import "app/domain"

type CategoriesRepository interface {
	FindAll(offset int, limit int) (categories []domain.Category, err error)
	Find(id int) (category domain.Category, err error)
	Create(form domain.Category) (err error)
	Update(form domain.Category) (err error)
	TotalCount() (count int64, err error)
}
