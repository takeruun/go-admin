package usecase

import (
	"app/domain"
	"app/usecase/repository"
)

type ProductsInteractor struct {
	Products repository.ProductsRepository
}

func (interactor *ProductsInteractor) Index(offset int, limit int, form domain.Product) (products []domain.Product, totalCount int64, err error) {
	products, err = interactor.Products.FindAll(offset, limit, form)
	if err != nil {
		return nil, 0, err
	}

	totalCount, err = interactor.Products.TotalCount(form)
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

func (interactor *ProductsInteractor) Show(id int) (product domain.Product, err error) {
	product, err = interactor.Products.FindById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (interactor *ProductsInteractor) Create(form domain.Product) (err error) {
	err = interactor.Products.Create(form)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *ProductsInteractor) Update(form domain.Product) (err error) {
	err = interactor.Products.Update(form)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *ProductsInteractor) Search(form domain.Product) (products []domain.Product, err error) {
	products, err = interactor.Products.Search(form)

	if err != nil {
		return nil, err
	}

	return products, nil
}
