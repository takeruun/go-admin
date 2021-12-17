package usecase

import (
	"app/domain"
	"app/usecase/repository"
	"math"
	"time"
)

type UsersInteractor struct {
	Users repository.UsersRepository
}

func (interactor *UsersInteractor) FindAll() (users []domain.User, err error) {
	users, err = interactor.Users.FindAll()

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (interactor *UsersInteractor) Index(offset int, limit int) (users []domain.User, totalCount int64, err error) {
	users, err = interactor.Users.FindAllAndOrderDate(offset, limit)
	if err != nil {
		return []domain.User{}, 0, err
	}

	for i := 0; i < len(users); i++ {
		data, err := interactor.Users.LastVisitDate(int(users[i].ID))
		if err != nil {
			return []domain.User{}, 0, err
		}

		if data.Format("2006-01-02") != "0001-01-01" {
			diff := time.Now().Sub(data)
			users[i].LastVistDates = ReturnIntPointer(int(math.Round(diff.Hours() / 24)))
		} else {
			users[i].LastVistDates = nil
		}
	}

	totalCount, err = interactor.Users.TotalCount()
	if err != nil {
		return []domain.User{}, 0, err
	}

	return users, totalCount, nil
}

func (interactor *UsersInteractor) Add(form domain.User) (user domain.User, err error) {
	user, err = interactor.Users.Create(form)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (interactor *UsersInteractor) Show(id int) (user domain.User, err error) {
	user, err = interactor.Users.Find(id)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (interactor *UsersInteractor) Update(form domain.User) (err error) {
	err = interactor.Users.Update(form)

	if err != nil {
		return err
	}

	return
}

func (interactor *UsersInteractor) Search(queries []map[string]string) (users []domain.User, err error) {

	users, err = interactor.Users.SearchByQuery(queries)

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}
