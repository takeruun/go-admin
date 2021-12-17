package repository

import (
	"app/domain"
	"time"
)

type UsersRepository interface {
	Find(id int) (user domain.User, err error)
	FindAll() (users []domain.User, err error)
	FindAllAndOrderDate(offset int, limit int) (users []domain.User, err error)
	Create(form domain.User) (user domain.User, err error)
	Update(form domain.User) (err error)
	SearchByQuery(queries []map[string]string) (users []domain.User, err error)
	TotalCount() (count int64, err error)
	LastVisitDate(userId int) (date time.Time, err error)
}
