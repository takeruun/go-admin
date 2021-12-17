package usecase

import (
	"app/domain"
	"app/usecase/repository"
)

type CategoriesInteractor struct {
	Categories repository.CategoriesRepository
}

func (interactor *CategoriesInteractor) Index(offset int, limit int) (categories []domain.Category, totalCount int64, err error) {
	categories, err = interactor.Categories.FindAll(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	totalCount, err = interactor.Categories.TotalCount()
	if err != nil {
		return nil, 0, err
	}

	return categories, totalCount, nil
}

func (interactor *CategoriesInteractor) Show(id int) (category domain.Category, err error) {
	category, err = interactor.Categories.Find(id)

	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (interactor *CategoriesInteractor) Create(form domain.Category) (err error) {
	err = interactor.Categories.Create(form)

	if err != nil {
		return err
	}

	return
}

func (interactor *CategoriesInteractor) Update(form domain.Category) (err error) {
	err = interactor.Categories.Update(form)

	if err != nil {
		return err
	}

	return
}

func (interactor *CategoriesInteractor) CategoriesJson() (categories []domain.Category, err error) {
	categories, err = interactor.Categories.FindAll(0, 0)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
