package database

import (
	"app/domain"
	"time"
)

type UsersRepository struct {
	DB DB
}

func (repo *UsersRepository) Find(id int) (user domain.User, err error) {
	db := repo.DB.Connect()

	result := db.First(&user, id)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return user, nil
}

func (repo *UsersRepository) FindAll() (users []domain.User, err error) {
	db := repo.DB.Connect()

	result := db.Find(&users)
	if result.Error != nil {
		return []domain.User{}, result.Error
	}

	return users, nil
}

func (repo *UsersRepository) FindAllAndOrderDate(offset int, limit int) (users []domain.User, err error) {
	db := repo.DB.Connect()

	result := db.Table("users").Select("users.*, o1.NextVisitDate, o2.PreviousVisitDate").
		Joins("LEFT JOIN (select min(date_of_visit) as NextVisitDate, user_id from orders where status = ? AND date_format(date_of_visit, '%Y-%m-%d') >= ? GROUP BY user_id) AS o1 ON o1.user_id = users.id", domain.Resarvation, time.Now().Format("2006-01-02")).
		Joins("LEFT JOIN (select max(date_of_exit) as PreviousVisitDate, user_id from orders where status = ? GROUP BY user_id) AS o2 ON o2.user_id = users.id", domain.Complete).
		Order("users.id").
		Offset(offset).Limit(limit).
		Scan(&users)
	if result.Error != nil {
		return []domain.User{}, result.Error
	}

	return users, nil
}

func (repo *UsersRepository) Create(form domain.User) (user domain.User, err error) {
	db := repo.DB.Connect()

	result := db.Create(&form)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	db.First(&user, form.ID)

	return user, nil
}

func (repo *UsersRepository) Update(form domain.User) (err error) {
	db := repo.DB.Connect()

	result := db.Updates(&form)

	if result.Error != nil {
		return result.Error
	}

	return
}

func (repo *UsersRepository) SearchByQuery(queries []map[string]string) (users []domain.User, err error) {
	db := repo.DB.Connect()

	for _, query := range queries {
		for k, v := range query {
			db = db.Where(k+"= ?", v)
		}
	}
	result := db.Find(&users)
	if result.Error != nil {
		return []domain.User{}, result.Error
	}

	return users, nil
}

func (repo *UsersRepository) TotalCount() (count int64, err error) {
	db := repo.DB.Connect()

	result := db.Model(domain.User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return
}

func (repo *UsersRepository) LastVisitDate(userId int) (data time.Time, err error) {
	db := repo.DB.Connect()

	db.Model(domain.User{}).Joins("inner join orders on orders.user_id = users.id").Select("MAX(date_of_visit)").Where("users.id = ?", userId).Group("users.id").Scan(&data)
	return data, nil
}
